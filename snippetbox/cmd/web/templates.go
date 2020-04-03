package main

import "nickherrig.com/snippetbox/pkg/models"


type templateData struct {
    Snippet  *models.Snippet
    Snippets []*models.Snippet
}
