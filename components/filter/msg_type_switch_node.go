/*
 * Copyright 2023 The RuleGo Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package filter

//规则链节点配置示例：
//{
//        "id": "s2",
//        "type": "msgTypeSwitch",
//        "name": "消息路由"
//      }
import (
	"github.com/rulego/rulego/api/types"
)

// KeyDefaultRelationType 如果无法匹配到节点通过Default关系查找节点
const KeyDefaultRelationType = "Default"

// KeyOtherRelationTypeName config配置的默认关系Properties key
const KeyOtherRelationTypeName = "defaultRelationType"

func init() {
	Registry.Add(&MsgTypeSwitchNode{})
}

// MsgTypeSwitchNode 根据传入的消息类型路由到一个或多个输出链
// 把消息通过类型发到正确的链,
type MsgTypeSwitchNode struct {
	defaultRelationType string
}

// Type 组件类型
func (x *MsgTypeSwitchNode) Type() string {
	return "msgTypeSwitch"
}

func (x *MsgTypeSwitchNode) New() types.Node {
	return &MsgTypeSwitchNode{}
}

// Init 初始化
func (x *MsgTypeSwitchNode) Init(ruleConfig types.Config, configuration types.Configuration) error {
	if v := ruleConfig.Properties.GetValue(KeyOtherRelationTypeName); v != "" {
		x.defaultRelationType = v
	} else {
		x.defaultRelationType = KeyDefaultRelationType
	}
	return nil
}

// OnMsg 处理消息
func (x *MsgTypeSwitchNode) OnMsg(ctx types.RuleContext, msg types.RuleMsg) {
	ctx.TellNextOrElse(msg, x.defaultRelationType, msg.Type)
}

// Destroy 销毁
func (x *MsgTypeSwitchNode) Destroy() {
}
