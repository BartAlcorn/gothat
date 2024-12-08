package embedded

import "embed"

//go:embed assets
var Assets embed.FS

type RemoteFile struct {
	Folder string
	Name   string
	URL    string
}

var Scripts = []RemoteFile{
	// {Folder: "others", Name: "theme-change@2.4.0.js", URL: "https://cdn.jsdelivr.net/npm/theme-change@2.4.0/index.js"},
	{Folder: "others", Name: "quill@2.0.2.js", URL: "https://cdn.jsdelivr.net/npm/quill@2.0.2/dist/quill.js"},
	// {Folder: "others", Name: "quill.js.map", URL: "https://cdn.jsdelivr.net/npm/quill@2.0.2/dist/quill.js.map"},
	// {Folder: "others", Name: "popperjs@2.1.8.js", URL: "https://unpkg.com/@popperjs/core@2.11.8/dist/umd/popper.min.js"},
	// {Folder: "others", Name: "popper.min.js.map", URL: "https://unpkg.com/@popperjs/core@2.11.8/dist/umd/popper.min.js.map"},
	// {Folder: "others", Name: "tippyjs@6.3.7.min.js", URL: "https://unpkg.com/tippy.js@6.3.7/dist/tippy-bundle.umd.min.js"},
	// {Folder: "others", Name: "tippy-bundle.umd.min.js.map", URL: "https://unpkg.com/tippy.js@6.3.7/dist/tippy-bundle.umd.min.js.map"},
	{Folder: "others", Name: "echarts@5.5.1.min.js", URL: "https://cdn.jsdelivr.net/npm/echarts@5.5.1/dist/echarts.min.js"},
	{Folder: "aggrid", Name: "ag-grid-community@32.3.3.min.js", URL: "https://cdn.jsdelivr.net/npm/ag-grid-community@32.3.3/dist/ag-grid-community.min.js"},
	{Folder: "aggrid", Name: "ag-grid-enterprise@32.3.3.min.js", URL: "https://cdn.jsdelivr.net/npm/ag-grid-enterprise@32.3.3/dist/ag-grid-enterprise.min.js"},
	{Folder: "htmx", Name: "htmx@2.0.3.min.js", URL: "https://unpkg.com/htmx.org@2.0.3/dist/htmx.min.js"},
	{Folder: "htmx", Name: "htmx-ext-sse@2.2.2.js", URL: "https://unpkg.com/htmx-ext-sse@2.2.2/sse.js"},
	{Folder: "htmx", Name: "htmx-ext-alpine-morph@2.0.0.js", URL: "https://unpkg.com/htmx-ext-alpine-morph@2.0.0/alpine-morph.js"},
	{Folder: "htmx", Name: "htmx-ext-include-vals@2.0.0.js", URL: "https://unpkg.com/htmx-ext-include-vals@2.0.0/include-vals.js"},
	// {Folder: "htmx", Name: "hyperscript.org@0.9.12.js", URL: "https://unpkg.com/hyperscript.org@0.9.12"},
	{Folder: "alpinejs", Name: "ui@3.14.5.min.js", URL: "https://unpkg.com/@alpinejs/ui@3.14.5/dist/cdn.min.js"},
	{Folder: "alpinejs", Name: "anchor@3.14.5.min.js", URL: "https://unpkg.com/@alpinejs/anchor@3.14.5/dist/cdn.min.js"},
	{Folder: "alpinejs", Name: "collapse@3.14.5.min.js", URL: "https://unpkg.com/@alpinejs/collapse@3.14.5/dist/cdn.min.js"},
	{Folder: "alpinejs", Name: "intersect@3.14.5.min.js", URL: "https://unpkg.com/@alpinejs/intersect@3.14.5/dist/cdn.min.js"},
	{Folder: "alpinejs", Name: "focus@3.14.5.min.js", URL: "https://unpkg.com/@alpinejs/focus@3.14.5/dist/cdn.min.js"},
	{Folder: "alpinejs", Name: "sort@3.14.5.min.js", URL: "https://unpkg.com/@alpinejs/sort@3.14.5/dist/cdn.min.js"},
	{Folder: "alpinejs", Name: "alpinejs@3.14.5.min.js", URL: "https://unpkg.com/alpinejs@3.14.5/dist/cdn.min.js"},
}

var StyleSheets = []RemoteFile{
	{Folder: "aggrid", Name: "ag-grid-community@32.3.3.min.css", URL: "https://cdn.jsdelivr.net/npm/ag-grid-community@32.3.3/styles/ag-grid.min.css"},
}
