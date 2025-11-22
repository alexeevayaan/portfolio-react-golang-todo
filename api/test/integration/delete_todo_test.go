//go:build integration

package test

func (s *Suite) Test_DeleteTodo() {
	id, err := s.todo.Create(ctx, title, description)
	s.NoError(err)

	t, err := s.todo.Get(ctx, id.String())
	s.NoError(err)

	s.Equal(title, t.Title)
	s.Equal(description, t.Description)

	err = s.todo.Delete(ctx, id.String())
	s.NoError(err)

	_, err = s.todo.Get(ctx, id.String())
	s.Contains(err.Error(), "not found")
}
