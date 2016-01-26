package network

import (
	"sync"

	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type Config struct {
	Log spec.Log `json:"-"`

	States map[string]spec.State `json:"states,omitempty"`
}

const (
	ObjectType spec.ObjectType = "network"
)

func DefaultConfig() Config {
	newStateConfig := state.DefaultConfig()
	newStateConfig.ObjectType = ObjectType

	newConfig := Config{
		Log: log.NewLog(log.DefaultConfig()),
		States: map[string]spec.State{
			"default": state.NewState(newStateConfig),
		},
	}

	return newConfig
}

func NewNetwork(config Config) spec.Network {
	newNetwork := &network{
		Config: config,
		Mutex:  sync.Mutex{},
	}

	return newNetwork
}

type network struct {
	Config

	Mutex sync.Mutex `json:"-"`
}

func (n *network) Copy() spec.Network {
	networkCopy := *n

	for key, state := range networkCopy.States {
		networkCopy.States[key] = state.Copy()
	}

	return &networkCopy
}

func (n *network) GetObjectID() spec.ObjectID {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	return n.States["default"].GetObjectID()
}

func (n *network) GetObjectType() spec.ObjectType {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	return n.States["default"].GetObjectType()
}

func (n *network) GetState(key string) (spec.State, error) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()

	if state, ok := n.States[key]; ok {
		return state, nil
	}

	return nil, maskAny(stateNotFoundError)
}

func (n *network) SetState(key string, state spec.State) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	n.States[key] = state
}

func (n *network) Trigger(imp spec.Impulse) (spec.Impulse, error) {
	// Track state.
	impState, err := imp.GetState("default")
	if err != nil {
		return nil, maskAny(err)
	}
	impState.SetNetwork(n)
	networkState, err := n.GetState("default")
	if err != nil {
		return nil, maskAny(err)
	}
	networkState.SetImpulse(imp)

	// Get first neuron.
	var neu spec.Neuron
	defaultNetworkState, err := n.GetState("default")
	if err != nil {
		return nil, maskAny(err)
	}
	neurons := defaultNetworkState.GetNeurons()
	if len(neurons) == 0 {
		initNetworkState, err := n.GetState("init")
		if err != nil {
			return nil, maskAny(err)
		}
		neuronID, err := initNetworkState.GetBytes("first-id")
		if err != nil {
			return nil, maskAny(err)
		}
		initFirstNeuron, err := initNetworkState.GetNeuronByID(spec.ObjectID(neuronID))
		if err != nil {
			return nil, maskAny(err)
		}
		neu = initFirstNeuron.Copy()

		// Track state.
		neuronState, err := neu.GetState("default")
		if err != nil {
			return nil, maskAny(err)
		}
		neuronState.SetNetwork(n)
		networkState, err = n.GetState("default")
		if err != nil {
			return nil, maskAny(err)
		}
		networkState.SetNeuron(neu)
		networkState.SetBytes("first-id", []byte(neu.GetObjectID()))
	} else {
		defaultNetworkState, err := n.GetState("default")
		if err != nil {
			return nil, maskAny(err)
		}
		neuronID, err := defaultNetworkState.GetBytes("first-id")
		if err != nil {
			return nil, maskAny(err)
		}
		neu, err = defaultNetworkState.GetNeuronByID(spec.ObjectID(neuronID))
		if err != nil {
			return nil, maskAny(err)
		}
	}

	imp, _, err = imp.WalkThrough(neu)
	if err != nil {
		return nil, maskAny(err)
	}

	return imp, nil
}