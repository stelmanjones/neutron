const DEBUG = Bun.env.DEBUG === undefined ? false : true;

await Bun.build({
	entrypoints: ["./api/index.ts"],
	outdir: "./packages",
	target: "browser",
	format: "esm",
	minify: !DEBUG,
	sourcemap: DEBUG ? "linked" : "none"
});
