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
	{Folder: "others", Name: "theme-change@2.4.0.js", URL: "https://cdn.jsdelivr.net/npm/theme-change@2.4.0/index.js"},
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
