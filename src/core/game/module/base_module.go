package module

// 定义继承模块

import (
	"game/api"
)

type SceneBaseModule struct {
	Scene api.IScene
}

func (s *SceneBaseModule) InitBaseModule(base interface{}) {
	s.Scene = base.(api.IScene)
}

func (s *SceneBaseModule) Release() {
	s.Scene = nil
}

type RoleBaseModule struct {
	Role api.IRole
}

func (r *RoleBaseModule) InitBaseModule(base interface{}) {
	r.Role = base.(api.IRole)
}

func (r *RoleBaseModule) Release() {
	r.Role = nil
}
