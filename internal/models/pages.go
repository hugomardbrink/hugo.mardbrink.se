package models

type Breadcrumb struct {
    Title string
    Route string
}

type Breadcrumbs = []Breadcrumb

type Page struct {
    Title string
    Description string
    Breadcrumbs Breadcrumbs
    IsHtmx bool
}

type HomePage struct {
    Page Page
}

type ProjectPage struct {
    Page Page
}


func NewPage(title string, description string, breadcrumbs Breadcrumbs, isHtmx bool) Page {
    return Page{
        Title: title,
        Description: description,
        Breadcrumbs: breadcrumbs,
        IsHtmx: isHtmx,
    }
}

func NewBreadcrumb(title string, route string) Breadcrumb {
    return Breadcrumb{
        Title: title,
        Route: route,
    }
}

func NewHomePage(page Page) HomePage {
    return HomePage{
        Page: page,
    }
}

func NewProjectPage(page Page) ProjectPage {
    return ProjectPage{
        Page: page,
    }
}

