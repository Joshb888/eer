/*
 Copyright 2016 Padduck, LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 	http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package pufferpanel

type Operation interface {
	Run(env Environment) error
}

type OperationFactory interface {
	Create(CreateOperation) (Operation, error)

	Key() string
}

type CreateOperation struct {
	OperationArgs        map[string]interface{}
	EnvironmentVariables map[string]string
	DataMap              map[string]interface{}
}
