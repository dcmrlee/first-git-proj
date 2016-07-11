# Copyright 2011 Google Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

6.out: basicblock.o lsg.o havlaklookfinder.o looptesterapp.o
	go tool link looptesterapp.o

basicblock.o: basicblock.go
	go tool compile  basicblock.go

lsg.o: lsg.go
	go tool compile lsg.go

havlaklookfinder.o: havlakloopfinder.go
	go tool compile havlakloopfinder.go

looptesterapp.o: looptesterapp.go
	go tool compile looptesterapp.go


run: 
	./6.out

clean:
	rm -f *.o ./a.out
