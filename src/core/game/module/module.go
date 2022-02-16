package module

// 定义功能模块共用接口操作

type IInitModule interface {
	Init()
}

type IDataModule interface {
	InitData()
	SaveData()
}

type ILoopModule interface {
	Loop()
}

type IBaseModule interface {
	InitBaseModule(base interface{})
	Release()
}

type IModuleMgr interface {
	Register(name string, module interface{})
	GetModules() map[string]interface{}
	StartModules()
}

type ModuleMgr struct {
	modules map[string]interface{}
}

func NewModuleMgr() *ModuleMgr {
	return new(ModuleMgr)
}

func (m *ModuleMgr) Register(name string, module interface{}) {
	if m.modules == nil {
		m.modules = make(map[string]interface{})
		m.modules[name] = module
		return
	}
	if _, ok := m.modules[name]; !ok {
		m.modules[name] = module
	}
}

func (m *ModuleMgr) GetModules() map[string]interface{} {
	return m.modules
}

func (m *ModuleMgr) StartModules() {
	for _, module := range m.modules {
		if initModule, ok := module.(IInitModule); ok {
			initModule.Init()
		}
		if loopModule, ok := module.(ILoopModule); ok {
			loopModule.Loop()
		}
	}
}
