module github.com/stelmanjones/neutron

replace github.com/stelmanjones/neutron/internal/webview => ./internal/webview

replace github.com/stelmanjones/neutron/internal => ./internal

replace github.com/stelmanjones/neutron/internal/webview/include => ./internal/webview/inlcude

replace github.com/stelmanjones/neutron/internal/mswebview2 => ./internal/mswebview2

replace github.com/stelmanjones/neutron/internal/mswebview2/include => ./internal/mswebview2/include

go 1.23.1

require github.com/charmbracelet/log v0.4.0

require (
	atomicgo.dev/keyboard v0.2.9 // indirect
	github.com/containerd/console v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/gookit/color v1.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stelmanjones/termtools v0.0.0-20241019125237-c3b6ecd35376 // indirect
	github.com/stelmanjones/termtools/usure v0.0.0-20241019125345-912843dc204b // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/charmbracelet/lipgloss v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/spf13/cobra v1.8.1
	github.com/spf13/viper v1.19.0
	github.com/stelmanjones/termtools/prompt v0.0.0-20241019125345-912843dc204b
	golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f // indirect
	golang.org/x/sys v0.20.0 // indirect
)
