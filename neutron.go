package neutron

import (
	_ "embed"
	"iter"
	"maps"
	"os"
	"path"
	"runtime"
	"slices"
	"sync"

	"github.com/charmbracelet/log"
	webview "github.com/stelmanjones/neutron/internal"
)

//go:embed packages/index.js
var jsAPI []byte

//go:embed template.html
var htmlTemplate string

type Page struct {
	URL       string
	StartPage bool
}

type Application struct {
	pages  map[string]Page
	view   webview.WebView
	mux    *sync.Mutex
	logger *log.Logger
	config *NeutronConfig
}

func New(debug bool) *Application {
	c := InitConfig()
	return &Application{
		view: webview.New(debug),
		mux:  &sync.Mutex{},
		logger: log.NewWithOptions(os.Stderr, log.Options{
			Level: func() log.Level {
				if c.Debug {
					return log.DebugLevel
				}
				return log.InfoLevel
			}(),
			TimeFormat:      "15:04:05",
			ReportTimestamp: true,
		}),
		config: c,
	}
}

// Run starts the application.
func (a *Application) Run() {
	defer a.view.Destroy()

	if len(a.pages) == 0 {
		a.view.SetHtml(htmlTemplate)
	} else {
		if !a.hasStartPage() {
			a.view.Navigate("file://" + slices.Collect(maps.Values(a.pages))[0].URL)
		} else {
			for _, page := range a.pages {
				if page.StartPage {
					a.view.Navigate("file://" + page.URL)
					break
				}
			}
		}
	}
	a.view.SetSize(a.config.Width, a.config.Height, webview.HintNone)
	a.view.SetTitle(a.config.Name)
	a.view.Init(string(jsAPI))
	a.view.Bind("debug", a.logger.Debug)
	a.view.Bind("info", a.logger.Info)
	a.view.Bind("warn", a.logger.Warn)
	a.view.Bind("error", a.logger.Error)
	a.view.Bind("setTitle", a.SetTitle)
	a.view.Bind("setSize", a.SetSize)
	// HACK Should i bind fatal here?
	//	a.view.Bind("fatal", a.fatal)
	a.view.Run()
}

func (a *Application) onConfigChange() {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.SetSize(a.config.Width, a.config.Height)
	a.SetTitle(a.config.Name)
}

// Pages returns all the registered pages.
func (a *Application) Pages() iter.Seq2[string, Page] {
	a.mux.Lock()
	defer a.mux.Unlock()
	return maps.All(a.pages)
}

// Navigate navigates to a page.
func (a *Application) Navigate(page string) error {
	for k, p := range a.Pages() {
		if k == page {
			a.view.Navigate("file://" + p.URL)
			return nil
		}
	}
	return ErrNoSuchPage
}

// NavigateURL navigates to a URL.
func (a *Application) NavigateURL(url string) {
	a.view.Navigate(url)
}

// SetTitle sets the title of the application.
func (a *Application) SetTitle(title string) {
	a.mux.Lock()

	a.config.Name = title
	a.view.SetTitle(a.config.Name)
	a.mux.Unlock()
}

// SetSize sets the size of the application.
func (a *Application) SetSize(w, h int) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.view.SetSize(w, h, webview.HintNone)
	a.config.Width = w
	a.config.Height = h
}

// Bind binds a callback function so that it will appear under the given name as a global JavaScript function. Internally it uses webview\_init(). Callback receives a request string and a user-provided argument pointer. Request string is a JSON array of all the arguments passed to the JavaScript function.
// f must be a function and must return either a value and an error or just an error.
//
//	type Result struct{
//	 Value int `json:"value"`
//	}
//
// func ExampleBind() {
//
//	 var count uint = 0
//
//	 a.Bind("increment", func() Result {
//	   count ++
//	 return Result{Value: count}
//	 })
//	}
func (a *Application) Bind(name string, f interface{}) error {
	err := a.view.Bind(name, f)
	if err != nil {
		return err
	}
	return nil
}

// pathToStartPage returns the path to the template page.
func pathToStartPage() string {
	_, currentFile, _, _ := runtime.Caller(0)
	dir := path.Dir(currentFile)
	return path.Join(dir, "template.html")
}

// hasStartPage returns true if the application has a start page.
func (a *Application) hasStartPage() bool {
	for _, p := range a.pages {
		if p.StartPage {
			return true
		}
	}
	return false
}
