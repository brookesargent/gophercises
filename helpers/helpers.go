package helpers

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

type PathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYAML(yml []byte) ([]PathUrl, error) {
	var pathUrls []PathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return pathUrls, err
	}
	return pathUrls, nil
}