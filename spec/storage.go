package spec

// Storage represents a persistency management object.
type Storage interface {
	// Get returns data associated with key. This is a simple key-value
	// relationship.
	Get(key string) (string, error)

	// GetElementsByScore looks up all elements associated with the given score.
	// To limit the number of returned elements, maxElements ca be used. Note
	// that the list has this scheme.
	//
	//     element1,element2,...
	//
	GetElementsByScore(key string, score float64, maxElements int) ([]string, error)

	// GetStringMap returns the hash map stored under the given key.
	GetStringMap(key string) (map[string]string, error)

	// GetHighestScoredElements searches a list that is ordered by their
	// element's score, and returns the elements and their corresponding scores,
	// where the highest scored element is the first in the returned list. Note
	// that the list has this scheme.
	//
	//     element1,score1,element2,score2,...
	//
	// Note that the resulting list will have the length of maxElements*2,
	// because the list contains the elements and their scores.
	//
	GetHighestScoredElements(key string, maxElements int) ([]string, error)

	Object

	// PushToSet pushes the given element to the set identified by the given key.
	// Note that a set is an unordered list containing distinct objects.
	PushToSet(key string, element string) error

	// RemoveFromSet removes the given element from the set identified by the
	// given key.
	RemoveFromSet(key string, element string) error

	// RemoveScoredElement removes the given element from the scored set under
	// key.
	RemoveScoredElement(key string, element string) error

	// Set stores the given key value pair. Once persisted, value can be
	// retrieved using Get.
	Set(key string, value string) error

	// SetElementByScore persists the given element in the weighted list
	// identified by key with respect to the given score.
	SetElementByScore(key, element string, score float64) error

	// SetStringMap stores the given stringMap under the given key.
	SetStringMap(key string, stringMap map[string]string) error

	// WalkScoredElements scans the scored set given by key and executes the
	// callback for each found element. Note that the walk might ignores the
	// score order.
	//
	// The walk is throttled. That means some amount of elements are fetched at
	// once from the storage. After all fetched elements iterated, the next batch
	// of elements is fetched to continue iteration, until the given set is
	// walked completely. The given closer can be used to end the walk
	// immediately.
	WalkScoredElements(key string, closer <-chan struct{}, cb func(element string, score float64) error) error

	// WalkSet scans the set given by key and executes the callback for each
	// found element.
	//
	// The walk is throttled. That means some amount of elements are fetched at
	// once from the storage. After all fetched elements iterated, the next batch
	// of elements is fetched to continue iteration, until the given set is
	// walked completely. The given closer can be used to end the walk
	// immediately.
	WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error
}
