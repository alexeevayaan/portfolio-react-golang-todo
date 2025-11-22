//go:build integration

package test

func (s *Suite) Test_GetTodo_Ok() {
	id, err := s.todo.Create(ctx, title, description)
	s.NoError(err)

	t, err := s.todo.Get(ctx, id.String())
	s.NoError(err)

	s.Equal(title, t.Title)
	s.Equal(description, t.Description)
}

func (s *Suite) Test_GetTodo_NotFound() {
	_, err := s.todo.Get(ctx, "4e38b195-9b3e-44e8-8db0-bf14f124ff2b")
	s.NotNil(err)

	s.Contains(err.Error(), "not found")
}
