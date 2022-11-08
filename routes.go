package main

func (s *server) routes() {
	s.router.Path("/schema/{id}").Methods("POST").Handler(s.handleSchemaUpload())
	s.router.Path("/schema/{id}").Methods("GET").Handler(s.handleSchemaDownload())
	s.router.Path("/validate/{id}").Methods("POST").Handler(s.handleSchemaValidate())
}
