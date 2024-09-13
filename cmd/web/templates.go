package main

import "snippetbox.jackmartin.jp/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
