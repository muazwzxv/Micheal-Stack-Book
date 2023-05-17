package domain

import "github.com/stackus/errors"

var (
	ErrStoreNameIsBlank     = errors.Wrap(errors.ErrBadRequest, "store name cannot be blank")
	ErrStoreLocationIsBlank = errors.Wrap(errors.ErrBadRequest, "store location cannot be blank")
  ErrStoreIsAlreadyParticipating = errors.Wrap(errors.ErrBadRequest, "store is already participating")
  ErrStoreIsAlreadyNotParticipating = errors.Wrap(errors.ErrBadRequest, "store is already not participating")
)

type Store struct {
	ID            string
	Name          string
	Location      string
	Participating bool
}

func CreateStore(
	id, name, location string,
) (*Store, error) {
	if name == "" {
		return nil, ErrStoreNameIsBlank
	}

	if location == "" {
		return nil, ErrStoreLocationIsBlank
	}

  store := &Store{
		ID:       id,
		Name:     name,
		Location: location,
	}

  return store, nil
}

func (s *Store) EnableParticipation() error {
  if s.Participating {
    return ErrStoreIsAlreadyParticipating
  }

  s.Participating = true

  return nil
}

func (s *Store) DisableParticipation() error {
  if !s.Participating {
    return ErrStoreIsAlreadyNotParticipating
  }

  s.Participating = false

  return nil
}
