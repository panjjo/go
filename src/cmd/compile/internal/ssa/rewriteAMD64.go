// autogenerated from gen/AMD64.rules: do not edit!
// generated with: cd gen; go run *.go
package ssa

func rewriteValueAMD64(v *Value, config *Config) bool {
	switch v.Op {
	case OpAMD64ADDQ:
		// match: (ADDQ x (MOVQconst [c]))
		// cond:
		// result: (ADDQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64MOVQconst {
				goto endacffd55e74ee0ff59ad58a18ddfc9973
			}
			c := v.Args[1].AuxInt
			v.Op = OpAMD64ADDQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto endacffd55e74ee0ff59ad58a18ddfc9973
	endacffd55e74ee0ff59ad58a18ddfc9973:
		;
		// match: (ADDQ (MOVQconst [c]) x)
		// cond:
		// result: (ADDQconst [c] x)
		{
			if v.Args[0].Op != OpAMD64MOVQconst {
				goto end7166f476d744ab7a51125959d3d3c7e2
			}
			c := v.Args[0].AuxInt
			x := v.Args[1]
			v.Op = OpAMD64ADDQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto end7166f476d744ab7a51125959d3d3c7e2
	end7166f476d744ab7a51125959d3d3c7e2:
		;
		// match: (ADDQ x (SHLQconst [3] y))
		// cond:
		// result: (LEAQ8 x y)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64SHLQconst {
				goto endc02313d35a0525d1d680cd58992e820d
			}
			if v.Args[1].AuxInt != 3 {
				goto endc02313d35a0525d1d680cd58992e820d
			}
			y := v.Args[1].Args[0]
			v.Op = OpAMD64LEAQ8
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto endc02313d35a0525d1d680cd58992e820d
	endc02313d35a0525d1d680cd58992e820d:
		;
	case OpAMD64ADDQconst:
		// match: (ADDQconst [c] (LEAQ8 [d] x y))
		// cond:
		// result: (LEAQ8 [addOff(c, d)] x y)
		{
			c := v.AuxInt
			if v.Args[0].Op != OpAMD64LEAQ8 {
				goto ende2cc681c9abf9913288803fb1b39e639
			}
			d := v.Args[0].AuxInt
			x := v.Args[0].Args[0]
			y := v.Args[0].Args[1]
			v.Op = OpAMD64LEAQ8
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = addOff(c, d)
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto ende2cc681c9abf9913288803fb1b39e639
	ende2cc681c9abf9913288803fb1b39e639:
		;
		// match: (ADDQconst [0] x)
		// cond:
		// result: (Copy x)
		{
			if v.AuxInt != 0 {
				goto end288952f259d4a1842f1e8d5c389b3f28
			}
			x := v.Args[0]
			v.Op = OpCopy
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto end288952f259d4a1842f1e8d5c389b3f28
	end288952f259d4a1842f1e8d5c389b3f28:
		;
	case OpAMD64ANDQ:
		// match: (ANDQ x (MOVQconst [c]))
		// cond:
		// result: (ANDQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64MOVQconst {
				goto endb98096e3bbb90933e39c88bf41c688a9
			}
			c := v.Args[1].AuxInt
			v.Op = OpAMD64ANDQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto endb98096e3bbb90933e39c88bf41c688a9
	endb98096e3bbb90933e39c88bf41c688a9:
		;
		// match: (ANDQ (MOVQconst [c]) x)
		// cond:
		// result: (ANDQconst [c] x)
		{
			if v.Args[0].Op != OpAMD64MOVQconst {
				goto endd313fd1897a0d2bc79eff70159a81b6b
			}
			c := v.Args[0].AuxInt
			x := v.Args[1]
			v.Op = OpAMD64ANDQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto endd313fd1897a0d2bc79eff70159a81b6b
	endd313fd1897a0d2bc79eff70159a81b6b:
		;
	case OpAMD64ANDQconst:
		// match: (ANDQconst [0] _)
		// cond:
		// result: (MOVQconst [0])
		{
			if v.AuxInt != 0 {
				goto endf2afa4d9d31c344d6638dcdced383cf1
			}
			v.Op = OpAMD64MOVQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = 0
			return true
		}
		goto endf2afa4d9d31c344d6638dcdced383cf1
	endf2afa4d9d31c344d6638dcdced383cf1:
		;
		// match: (ANDQconst [-1] x)
		// cond:
		// result: (Copy x)
		{
			if v.AuxInt != -1 {
				goto end646afc7b328db89ad16ebfa156ae26e5
			}
			x := v.Args[0]
			v.Op = OpCopy
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto end646afc7b328db89ad16ebfa156ae26e5
	end646afc7b328db89ad16ebfa156ae26e5:
		;
	case OpAdd:
		// match: (Add <t> x y)
		// cond: (is64BitInt(t) || isPtr(t))
		// result: (ADDQ x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t) || isPtr(t)) {
				goto endf031c523d7dd08e4b8e7010a94cd94c9
			}
			v.Op = OpAMD64ADDQ
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto endf031c523d7dd08e4b8e7010a94cd94c9
	endf031c523d7dd08e4b8e7010a94cd94c9:
		;
		// match: (Add <t> x y)
		// cond: is32BitInt(t)
		// result: (ADDL x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is32BitInt(t)) {
				goto end35a02a1587264e40cf1055856ff8445a
			}
			v.Op = OpAMD64ADDL
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto end35a02a1587264e40cf1055856ff8445a
	end35a02a1587264e40cf1055856ff8445a:
		;
	case OpAMD64CMOVQCC:
		// match: (CMOVQCC (CMPQconst [c] (MOVQconst [d])) _ x)
		// cond: inBounds(d, c)
		// result: (Copy x)
		{
			if v.Args[0].Op != OpAMD64CMPQconst {
				goto endd5357f3fd5516dcc859c8c5b3c9efaa4
			}
			c := v.Args[0].AuxInt
			if v.Args[0].Args[0].Op != OpAMD64MOVQconst {
				goto endd5357f3fd5516dcc859c8c5b3c9efaa4
			}
			d := v.Args[0].Args[0].AuxInt
			x := v.Args[2]
			if !(inBounds(d, c)) {
				goto endd5357f3fd5516dcc859c8c5b3c9efaa4
			}
			v.Op = OpCopy
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto endd5357f3fd5516dcc859c8c5b3c9efaa4
	endd5357f3fd5516dcc859c8c5b3c9efaa4:
		;
		// match: (CMOVQCC (CMPQconst [c] (MOVQconst [d])) x _)
		// cond: !inBounds(d, c)
		// result: (Copy x)
		{
			if v.Args[0].Op != OpAMD64CMPQconst {
				goto end6ad8b1758415a9afe758272b34970d5d
			}
			c := v.Args[0].AuxInt
			if v.Args[0].Args[0].Op != OpAMD64MOVQconst {
				goto end6ad8b1758415a9afe758272b34970d5d
			}
			d := v.Args[0].Args[0].AuxInt
			x := v.Args[1]
			if !(!inBounds(d, c)) {
				goto end6ad8b1758415a9afe758272b34970d5d
			}
			v.Op = OpCopy
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto end6ad8b1758415a9afe758272b34970d5d
	end6ad8b1758415a9afe758272b34970d5d:
		;
	case OpAMD64CMPQ:
		// match: (CMPQ x (MOVQconst [c]))
		// cond:
		// result: (CMPQconst x [c])
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64MOVQconst {
				goto end32ef1328af280ac18fa8045a3502dae9
			}
			c := v.Args[1].AuxInt
			v.Op = OpAMD64CMPQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AuxInt = c
			return true
		}
		goto end32ef1328af280ac18fa8045a3502dae9
	end32ef1328af280ac18fa8045a3502dae9:
		;
		// match: (CMPQ (MOVQconst [c]) x)
		// cond:
		// result: (InvertFlags (CMPQconst <TypeFlags> x [c]))
		{
			if v.Args[0].Op != OpAMD64MOVQconst {
				goto endf8ca12fe79290bc82b11cfa463bc9413
			}
			c := v.Args[0].AuxInt
			x := v.Args[1]
			v.Op = OpAMD64InvertFlags
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAMD64CMPQconst, TypeInvalid)
			v0.Type = TypeFlags
			v0.AddArg(x)
			v0.AuxInt = c
			v.AddArg(v0)
			return true
		}
		goto endf8ca12fe79290bc82b11cfa463bc9413
	endf8ca12fe79290bc82b11cfa463bc9413:
		;
	case OpClosureCall:
		// match: (ClosureCall entry closure mem)
		// cond:
		// result: (CALLclosure entry closure mem)
		{
			entry := v.Args[0]
			closure := v.Args[1]
			mem := v.Args[2]
			v.Op = OpAMD64CALLclosure
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(entry)
			v.AddArg(closure)
			v.AddArg(mem)
			return true
		}
		goto endee26da781e813a3c602ccb4f7ade98c7
	endee26da781e813a3c602ccb4f7ade98c7:
		;
	case OpConst:
		// match: (Const <t> [val])
		// cond: is64BitInt(t)
		// result: (MOVQconst [val])
		{
			t := v.Type
			val := v.AuxInt
			if !(is64BitInt(t)) {
				goto end7f5c5b34093fbc6860524cb803ee51bf
			}
			v.Op = OpAMD64MOVQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = val
			return true
		}
		goto end7f5c5b34093fbc6860524cb803ee51bf
	end7f5c5b34093fbc6860524cb803ee51bf:
		;
	case OpGlobal:
		// match: (Global {sym})
		// cond:
		// result: (LEAQglobal {sym})
		{
			sym := v.Aux
			v.Op = OpAMD64LEAQglobal
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Aux = sym
			return true
		}
		goto end8f47b6f351fecaeded45abbe5c2beec0
	end8f47b6f351fecaeded45abbe5c2beec0:
		;
	case OpIsInBounds:
		// match: (IsInBounds idx len)
		// cond:
		// result: (SETB (CMPQ <TypeFlags> idx len))
		{
			idx := v.Args[0]
			len := v.Args[1]
			v.Op = OpAMD64SETB
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAMD64CMPQ, TypeInvalid)
			v0.Type = TypeFlags
			v0.AddArg(idx)
			v0.AddArg(len)
			v.AddArg(v0)
			return true
		}
		goto endb51d371171154c0f1613b687757e0576
	endb51d371171154c0f1613b687757e0576:
		;
	case OpIsNonNil:
		// match: (IsNonNil p)
		// cond:
		// result: (SETNE (TESTQ <TypeFlags> p p))
		{
			p := v.Args[0]
			v.Op = OpAMD64SETNE
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAMD64TESTQ, TypeInvalid)
			v0.Type = TypeFlags
			v0.AddArg(p)
			v0.AddArg(p)
			v.AddArg(v0)
			return true
		}
		goto endff508c3726edfb573abc6128c177e76c
	endff508c3726edfb573abc6128c177e76c:
		;
	case OpLess:
		// match: (Less x y)
		// cond: is64BitInt(v.Args[0].Type) && isSigned(v.Args[0].Type)
		// result: (SETL (CMPQ <TypeFlags> x y))
		{
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(v.Args[0].Type) && isSigned(v.Args[0].Type)) {
				goto endcecf13a952d4c6c2383561c7d68a3cf9
			}
			v.Op = OpAMD64SETL
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAMD64CMPQ, TypeInvalid)
			v0.Type = TypeFlags
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		goto endcecf13a952d4c6c2383561c7d68a3cf9
	endcecf13a952d4c6c2383561c7d68a3cf9:
		;
	case OpLoad:
		// match: (Load <t> ptr mem)
		// cond: t.IsBoolean()
		// result: (MOVBload ptr mem)
		{
			t := v.Type
			ptr := v.Args[0]
			mem := v.Args[1]
			if !(t.IsBoolean()) {
				goto endc119e594c7f8e8ce5ff97c00b501dba0
			}
			v.Op = OpAMD64MOVBload
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto endc119e594c7f8e8ce5ff97c00b501dba0
	endc119e594c7f8e8ce5ff97c00b501dba0:
		;
		// match: (Load <t> ptr mem)
		// cond: (is64BitInt(t) || isPtr(t))
		// result: (MOVQload ptr mem)
		{
			t := v.Type
			ptr := v.Args[0]
			mem := v.Args[1]
			if !(is64BitInt(t) || isPtr(t)) {
				goto end7c4c53acf57ebc5f03273652ba1d5934
			}
			v.Op = OpAMD64MOVQload
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto end7c4c53acf57ebc5f03273652ba1d5934
	end7c4c53acf57ebc5f03273652ba1d5934:
		;
	case OpLsh:
		// match: (Lsh <t> x y)
		// cond: is64BitInt(t)
		// result: (ANDQ (SHLQ <t> x y) (SBBQcarrymask <t> (CMPQconst <TypeFlags> [64] y)))
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t)) {
				goto end5d9e2211940fbc82536685578cf37d08
			}
			v.Op = OpAMD64ANDQ
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAMD64SHLQ, TypeInvalid)
			v0.Type = t
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpAMD64SBBQcarrymask, TypeInvalid)
			v1.Type = t
			v2 := v.Block.NewValue0(v.Line, OpAMD64CMPQconst, TypeInvalid)
			v2.Type = TypeFlags
			v2.AuxInt = 64
			v2.AddArg(y)
			v1.AddArg(v2)
			v.AddArg(v1)
			return true
		}
		goto end5d9e2211940fbc82536685578cf37d08
	end5d9e2211940fbc82536685578cf37d08:
		;
	case OpAMD64MOVQload:
		// match: (MOVQload [off1] (ADDQconst [off2] ptr) mem)
		// cond:
		// result: (MOVQload [addOff(off1, off2)] ptr mem)
		{
			off1 := v.AuxInt
			if v.Args[0].Op != OpAMD64ADDQconst {
				goto end843d29b538c4483b432b632e5666d6e3
			}
			off2 := v.Args[0].AuxInt
			ptr := v.Args[0].Args[0]
			mem := v.Args[1]
			v.Op = OpAMD64MOVQload
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(mem)
			return true
		}
		goto end843d29b538c4483b432b632e5666d6e3
	end843d29b538c4483b432b632e5666d6e3:
		;
		// match: (MOVQload [off1] (LEAQ8 [off2] ptr idx) mem)
		// cond:
		// result: (MOVQloadidx8 [addOff(off1, off2)] ptr idx mem)
		{
			off1 := v.AuxInt
			if v.Args[0].Op != OpAMD64LEAQ8 {
				goto end02f5ad148292c46463e7c20d3b821735
			}
			off2 := v.Args[0].AuxInt
			ptr := v.Args[0].Args[0]
			idx := v.Args[0].Args[1]
			mem := v.Args[1]
			v.Op = OpAMD64MOVQloadidx8
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		goto end02f5ad148292c46463e7c20d3b821735
	end02f5ad148292c46463e7c20d3b821735:
		;
	case OpAMD64MOVQloadidx8:
		// match: (MOVQloadidx8 [off1] (ADDQconst [off2] ptr) idx mem)
		// cond:
		// result: (MOVQloadidx8 [addOff(off1, off2)] ptr idx mem)
		{
			off1 := v.AuxInt
			if v.Args[0].Op != OpAMD64ADDQconst {
				goto ende81e44bcfb11f90916ccb440c590121f
			}
			off2 := v.Args[0].AuxInt
			ptr := v.Args[0].Args[0]
			idx := v.Args[1]
			mem := v.Args[2]
			v.Op = OpAMD64MOVQloadidx8
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(mem)
			return true
		}
		goto ende81e44bcfb11f90916ccb440c590121f
	ende81e44bcfb11f90916ccb440c590121f:
		;
	case OpAMD64MOVQstore:
		// match: (MOVQstore [off1] (ADDQconst [off2] ptr) val mem)
		// cond:
		// result: (MOVQstore [addOff(off1, off2)] ptr val mem)
		{
			off1 := v.AuxInt
			if v.Args[0].Op != OpAMD64ADDQconst {
				goto end2108c693a43c79aed10b9246c39c80aa
			}
			off2 := v.Args[0].AuxInt
			ptr := v.Args[0].Args[0]
			val := v.Args[1]
			mem := v.Args[2]
			v.Op = OpAMD64MOVQstore
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		goto end2108c693a43c79aed10b9246c39c80aa
	end2108c693a43c79aed10b9246c39c80aa:
		;
		// match: (MOVQstore [off1] (LEAQ8 [off2] ptr idx) val mem)
		// cond:
		// result: (MOVQstoreidx8 [addOff(off1, off2)] ptr idx val mem)
		{
			off1 := v.AuxInt
			if v.Args[0].Op != OpAMD64LEAQ8 {
				goto endce1db8c8d37c8397c500a2068a65c215
			}
			off2 := v.Args[0].AuxInt
			ptr := v.Args[0].Args[0]
			idx := v.Args[0].Args[1]
			val := v.Args[1]
			mem := v.Args[2]
			v.Op = OpAMD64MOVQstoreidx8
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		goto endce1db8c8d37c8397c500a2068a65c215
	endce1db8c8d37c8397c500a2068a65c215:
		;
	case OpAMD64MOVQstoreidx8:
		// match: (MOVQstoreidx8 [off1] (ADDQconst [off2] ptr) idx val mem)
		// cond:
		// result: (MOVQstoreidx8 [addOff(off1, off2)] ptr idx val mem)
		{
			off1 := v.AuxInt
			if v.Args[0].Op != OpAMD64ADDQconst {
				goto end01c970657b0fdefeab82458c15022163
			}
			off2 := v.Args[0].AuxInt
			ptr := v.Args[0].Args[0]
			idx := v.Args[1]
			val := v.Args[2]
			mem := v.Args[3]
			v.Op = OpAMD64MOVQstoreidx8
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = addOff(off1, off2)
			v.AddArg(ptr)
			v.AddArg(idx)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		goto end01c970657b0fdefeab82458c15022163
	end01c970657b0fdefeab82458c15022163:
		;
	case OpAMD64MULQ:
		// match: (MULQ x (MOVQconst [c]))
		// cond: c == int64(int32(c))
		// result: (MULQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64MOVQconst {
				goto end680a32a37babfff4bfa7d23be592a131
			}
			c := v.Args[1].AuxInt
			if !(c == int64(int32(c))) {
				goto end680a32a37babfff4bfa7d23be592a131
			}
			v.Op = OpAMD64MULQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto end680a32a37babfff4bfa7d23be592a131
	end680a32a37babfff4bfa7d23be592a131:
		;
		// match: (MULQ (MOVQconst [c]) x)
		// cond:
		// result: (MULQconst [c] x)
		{
			if v.Args[0].Op != OpAMD64MOVQconst {
				goto endc6e18d6968175d6e58eafa6dcf40c1b8
			}
			c := v.Args[0].AuxInt
			x := v.Args[1]
			v.Op = OpAMD64MULQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto endc6e18d6968175d6e58eafa6dcf40c1b8
	endc6e18d6968175d6e58eafa6dcf40c1b8:
		;
	case OpAMD64MULQconst:
		// match: (MULQconst [8] x)
		// cond:
		// result: (SHLQconst [3] x)
		{
			if v.AuxInt != 8 {
				goto ende8d313a52a134fb2e1c0beb54ea599fd
			}
			x := v.Args[0]
			v.Op = OpAMD64SHLQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = 3
			v.AddArg(x)
			return true
		}
		goto ende8d313a52a134fb2e1c0beb54ea599fd
	ende8d313a52a134fb2e1c0beb54ea599fd:
		;
		// match: (MULQconst [64] x)
		// cond:
		// result: (SHLQconst [5] x)
		{
			if v.AuxInt != 64 {
				goto end75c0c250c703f89e6c43d718dd5ea3c0
			}
			x := v.Args[0]
			v.Op = OpAMD64SHLQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = 5
			v.AddArg(x)
			return true
		}
		goto end75c0c250c703f89e6c43d718dd5ea3c0
	end75c0c250c703f89e6c43d718dd5ea3c0:
		;
	case OpMove:
		// match: (Move [size] dst src mem)
		// cond:
		// result: (REPMOVSB dst src (Const <TypeUInt64> [size]) mem)
		{
			size := v.AuxInt
			dst := v.Args[0]
			src := v.Args[1]
			mem := v.Args[2]
			v.Op = OpAMD64REPMOVSB
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(dst)
			v.AddArg(src)
			v0 := v.Block.NewValue0(v.Line, OpConst, TypeInvalid)
			v0.Type = TypeUInt64
			v0.AuxInt = size
			v.AddArg(v0)
			v.AddArg(mem)
			return true
		}
		goto end1b2d226705fd31dbbe74e3286af178ea
	end1b2d226705fd31dbbe74e3286af178ea:
		;
	case OpMul:
		// match: (Mul <t> x y)
		// cond: is64BitInt(t)
		// result: (MULQ x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t)) {
				goto endfab0d598f376ecba45a22587d50f7aff
			}
			v.Op = OpAMD64MULQ
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto endfab0d598f376ecba45a22587d50f7aff
	endfab0d598f376ecba45a22587d50f7aff:
		;
	case OpOffPtr:
		// match: (OffPtr [off] ptr)
		// cond:
		// result: (ADDQconst [off] ptr)
		{
			off := v.AuxInt
			ptr := v.Args[0]
			v.Op = OpAMD64ADDQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = off
			v.AddArg(ptr)
			return true
		}
		goto end0429f947ee7ac49ff45a243e461a5290
	end0429f947ee7ac49ff45a243e461a5290:
		;
	case OpRsh:
		// match: (Rsh <t> x y)
		// cond: is64BitInt(t) && !t.IsSigned()
		// result: (ANDQ (SHRQ <t> x y) (SBBQcarrymask <t> (CMPQconst <TypeFlags> [64] y)))
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t) && !t.IsSigned()) {
				goto ende3e068773b8e6def1eaedb4f404ca6e5
			}
			v.Op = OpAMD64ANDQ
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAMD64SHRQ, TypeInvalid)
			v0.Type = t
			v0.AddArg(x)
			v0.AddArg(y)
			v.AddArg(v0)
			v1 := v.Block.NewValue0(v.Line, OpAMD64SBBQcarrymask, TypeInvalid)
			v1.Type = t
			v2 := v.Block.NewValue0(v.Line, OpAMD64CMPQconst, TypeInvalid)
			v2.Type = TypeFlags
			v2.AuxInt = 64
			v2.AddArg(y)
			v1.AddArg(v2)
			v.AddArg(v1)
			return true
		}
		goto ende3e068773b8e6def1eaedb4f404ca6e5
	ende3e068773b8e6def1eaedb4f404ca6e5:
		;
		// match: (Rsh <t> x y)
		// cond: is64BitInt(t) && t.IsSigned()
		// result: (SARQ <t> x (CMOVQCC <t> 			(CMPQconst <TypeFlags> [64] y) 			(Const <t> [63]) 			y))
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t) && t.IsSigned()) {
				goto end901ea4851cd5d2277a1ca1bee8f69d59
			}
			v.Op = OpAMD64SARQ
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.Type = t
			v.AddArg(x)
			v0 := v.Block.NewValue0(v.Line, OpAMD64CMOVQCC, TypeInvalid)
			v0.Type = t
			v1 := v.Block.NewValue0(v.Line, OpAMD64CMPQconst, TypeInvalid)
			v1.Type = TypeFlags
			v1.AuxInt = 64
			v1.AddArg(y)
			v0.AddArg(v1)
			v2 := v.Block.NewValue0(v.Line, OpConst, TypeInvalid)
			v2.Type = t
			v2.AuxInt = 63
			v0.AddArg(v2)
			v0.AddArg(y)
			v.AddArg(v0)
			return true
		}
		goto end901ea4851cd5d2277a1ca1bee8f69d59
	end901ea4851cd5d2277a1ca1bee8f69d59:
		;
	case OpAMD64SARQ:
		// match: (SARQ x (MOVQconst [c]))
		// cond:
		// result: (SARQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64MOVQconst {
				goto end031712b4008075e25a5827dcb8dd3ebb
			}
			c := v.Args[1].AuxInt
			v.Op = OpAMD64SARQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto end031712b4008075e25a5827dcb8dd3ebb
	end031712b4008075e25a5827dcb8dd3ebb:
		;
	case OpAMD64SBBQcarrymask:
		// match: (SBBQcarrymask (CMPQconst [c] (MOVQconst [d])))
		// cond: inBounds(d, c)
		// result: (Const [-1])
		{
			if v.Args[0].Op != OpAMD64CMPQconst {
				goto endf67d323ecef000dbcd15d7e031c3475e
			}
			c := v.Args[0].AuxInt
			if v.Args[0].Args[0].Op != OpAMD64MOVQconst {
				goto endf67d323ecef000dbcd15d7e031c3475e
			}
			d := v.Args[0].Args[0].AuxInt
			if !(inBounds(d, c)) {
				goto endf67d323ecef000dbcd15d7e031c3475e
			}
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = -1
			return true
		}
		goto endf67d323ecef000dbcd15d7e031c3475e
	endf67d323ecef000dbcd15d7e031c3475e:
		;
		// match: (SBBQcarrymask (CMPQconst [c] (MOVQconst [d])))
		// cond: !inBounds(d, c)
		// result: (Const [0])
		{
			if v.Args[0].Op != OpAMD64CMPQconst {
				goto end4157ddea9c4f71bfabfd6fa50e1208ed
			}
			c := v.Args[0].AuxInt
			if v.Args[0].Args[0].Op != OpAMD64MOVQconst {
				goto end4157ddea9c4f71bfabfd6fa50e1208ed
			}
			d := v.Args[0].Args[0].AuxInt
			if !(!inBounds(d, c)) {
				goto end4157ddea9c4f71bfabfd6fa50e1208ed
			}
			v.Op = OpConst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = 0
			return true
		}
		goto end4157ddea9c4f71bfabfd6fa50e1208ed
	end4157ddea9c4f71bfabfd6fa50e1208ed:
		;
	case OpAMD64SETG:
		// match: (SETG (InvertFlags x))
		// cond:
		// result: (SETL x)
		{
			if v.Args[0].Op != OpAMD64InvertFlags {
				goto endf7586738694c9cd0b74ae28bbadb649f
			}
			x := v.Args[0].Args[0]
			v.Op = OpAMD64SETL
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto endf7586738694c9cd0b74ae28bbadb649f
	endf7586738694c9cd0b74ae28bbadb649f:
		;
	case OpAMD64SETL:
		// match: (SETL (InvertFlags x))
		// cond:
		// result: (SETG x)
		{
			if v.Args[0].Op != OpAMD64InvertFlags {
				goto ende33160cd86b9d4d3b77e02fb4658d5d3
			}
			x := v.Args[0].Args[0]
			v.Op = OpAMD64SETG
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			return true
		}
		goto ende33160cd86b9d4d3b77e02fb4658d5d3
	ende33160cd86b9d4d3b77e02fb4658d5d3:
		;
	case OpAMD64SHLQ:
		// match: (SHLQ x (MOVQconst [c]))
		// cond:
		// result: (SHLQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64MOVQconst {
				goto endcca412bead06dc3d56ef034a82d184d6
			}
			c := v.Args[1].AuxInt
			v.Op = OpAMD64SHLQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto endcca412bead06dc3d56ef034a82d184d6
	endcca412bead06dc3d56ef034a82d184d6:
		;
	case OpAMD64SHRQ:
		// match: (SHRQ x (MOVQconst [c]))
		// cond:
		// result: (SHRQconst [c] x)
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64MOVQconst {
				goto endbb0d3a04dd2b810cb3dbdf7ef665f22b
			}
			c := v.Args[1].AuxInt
			v.Op = OpAMD64SHRQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = c
			v.AddArg(x)
			return true
		}
		goto endbb0d3a04dd2b810cb3dbdf7ef665f22b
	endbb0d3a04dd2b810cb3dbdf7ef665f22b:
		;
	case OpAMD64SUBQ:
		// match: (SUBQ x (MOVQconst [c]))
		// cond:
		// result: (SUBQconst x [c])
		{
			x := v.Args[0]
			if v.Args[1].Op != OpAMD64MOVQconst {
				goto end5a74a63bd9ad15437717c6df3b25eebb
			}
			c := v.Args[1].AuxInt
			v.Op = OpAMD64SUBQconst
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AuxInt = c
			return true
		}
		goto end5a74a63bd9ad15437717c6df3b25eebb
	end5a74a63bd9ad15437717c6df3b25eebb:
		;
		// match: (SUBQ <t> (MOVQconst [c]) x)
		// cond:
		// result: (NEGQ (SUBQconst <t> x [c]))
		{
			t := v.Type
			if v.Args[0].Op != OpAMD64MOVQconst {
				goto end78e66b6fc298684ff4ac8aec5ce873c9
			}
			c := v.Args[0].AuxInt
			x := v.Args[1]
			v.Op = OpAMD64NEGQ
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v0 := v.Block.NewValue0(v.Line, OpAMD64SUBQconst, TypeInvalid)
			v0.Type = t
			v0.AddArg(x)
			v0.AuxInt = c
			v.AddArg(v0)
			return true
		}
		goto end78e66b6fc298684ff4ac8aec5ce873c9
	end78e66b6fc298684ff4ac8aec5ce873c9:
		;
	case OpStaticCall:
		// match: (StaticCall [target] mem)
		// cond:
		// result: (CALLstatic [target] mem)
		{
			target := v.AuxInt
			mem := v.Args[0]
			v.Op = OpAMD64CALLstatic
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AuxInt = target
			v.AddArg(mem)
			return true
		}
		goto endcf02eb60d90086f6c42bfdc5842b145d
	endcf02eb60d90086f6c42bfdc5842b145d:
		;
	case OpStore:
		// match: (Store ptr val mem)
		// cond: (is64BitInt(val.Type) || isPtr(val.Type))
		// result: (MOVQstore ptr val mem)
		{
			ptr := v.Args[0]
			val := v.Args[1]
			mem := v.Args[2]
			if !(is64BitInt(val.Type) || isPtr(val.Type)) {
				goto endbaeb60123806948cd2433605820d5af1
			}
			v.Op = OpAMD64MOVQstore
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(ptr)
			v.AddArg(val)
			v.AddArg(mem)
			return true
		}
		goto endbaeb60123806948cd2433605820d5af1
	endbaeb60123806948cd2433605820d5af1:
		;
	case OpSub:
		// match: (Sub <t> x y)
		// cond: is64BitInt(t)
		// result: (SUBQ x y)
		{
			t := v.Type
			x := v.Args[0]
			y := v.Args[1]
			if !(is64BitInt(t)) {
				goto ende6ef29f885a8ecf3058212bb95917323
			}
			v.Op = OpAMD64SUBQ
			v.AuxInt = 0
			v.Aux = nil
			v.resetArgs()
			v.AddArg(x)
			v.AddArg(y)
			return true
		}
		goto ende6ef29f885a8ecf3058212bb95917323
	ende6ef29f885a8ecf3058212bb95917323:
	}
	return false
}
func rewriteBlockAMD64(b *Block) bool {
	switch b.Kind {
	case BlockAMD64EQ:
		// match: (EQ (InvertFlags cmp) yes no)
		// cond:
		// result: (EQ cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto end6b8e9afc73b1c4d528f31a60d2575fae
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64EQ
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end6b8e9afc73b1c4d528f31a60d2575fae
	end6b8e9afc73b1c4d528f31a60d2575fae:
		;
	case BlockAMD64GE:
		// match: (GE (InvertFlags cmp) yes no)
		// cond:
		// result: (LE cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto end0610f000a6988ee8310307ec2ea138f8
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64LE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end0610f000a6988ee8310307ec2ea138f8
	end0610f000a6988ee8310307ec2ea138f8:
		;
	case BlockAMD64GT:
		// match: (GT (InvertFlags cmp) yes no)
		// cond:
		// result: (LT cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto endf60c0660b6a8aa9565c97fc87f04eb34
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64LT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto endf60c0660b6a8aa9565c97fc87f04eb34
	endf60c0660b6a8aa9565c97fc87f04eb34:
		;
	case BlockIf:
		// match: (If (SETL cmp) yes no)
		// cond:
		// result: (LT cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64SETL {
				goto ende4d36879bb8e1bd8facaa8c91ba99dcc
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64LT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto ende4d36879bb8e1bd8facaa8c91ba99dcc
	ende4d36879bb8e1bd8facaa8c91ba99dcc:
		;
		// match: (If (SETNE cmp) yes no)
		// cond:
		// result: (NE cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64SETNE {
				goto end5ff1403aaf7b543bc454177ab584e4f5
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64NE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end5ff1403aaf7b543bc454177ab584e4f5
	end5ff1403aaf7b543bc454177ab584e4f5:
		;
		// match: (If (SETB cmp) yes no)
		// cond:
		// result: (ULT cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64SETB {
				goto end04935012db9defeafceef8175f803ea2
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64ULT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end04935012db9defeafceef8175f803ea2
	end04935012db9defeafceef8175f803ea2:
		;
		// match: (If cond yes no)
		// cond: cond.Op == OpAMD64MOVBload
		// result: (NE (TESTB <TypeFlags> cond cond) yes no)
		{
			v := b.Control
			cond := v
			yes := b.Succs[0]
			no := b.Succs[1]
			if !(cond.Op == OpAMD64MOVBload) {
				goto end7e22019fb0effc80f85c05ea30bdb5d9
			}
			b.Kind = BlockAMD64NE
			v0 := v.Block.NewValue0(v.Line, OpAMD64TESTB, TypeInvalid)
			v0.Type = TypeFlags
			v0.AddArg(cond)
			v0.AddArg(cond)
			b.Control = v0
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end7e22019fb0effc80f85c05ea30bdb5d9
	end7e22019fb0effc80f85c05ea30bdb5d9:
		;
	case BlockAMD64LE:
		// match: (LE (InvertFlags cmp) yes no)
		// cond:
		// result: (GE cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto end0d49d7d087fe7578e8015cf13dae37e3
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64GE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end0d49d7d087fe7578e8015cf13dae37e3
	end0d49d7d087fe7578e8015cf13dae37e3:
		;
	case BlockAMD64LT:
		// match: (LT (InvertFlags cmp) yes no)
		// cond:
		// result: (GT cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto end6a408cde0fee0ae7b7da0443c8d902bf
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64GT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end6a408cde0fee0ae7b7da0443c8d902bf
	end6a408cde0fee0ae7b7da0443c8d902bf:
		;
	case BlockAMD64NE:
		// match: (NE (InvertFlags cmp) yes no)
		// cond:
		// result: (NE cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto end713001aba794e50b582fbff930e110af
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64NE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end713001aba794e50b582fbff930e110af
	end713001aba794e50b582fbff930e110af:
		;
	case BlockAMD64UGE:
		// match: (UGE (InvertFlags cmp) yes no)
		// cond:
		// result: (ULE cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto ende3e4ddc183ca1a46598b11c2d0d13966
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64ULE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto ende3e4ddc183ca1a46598b11c2d0d13966
	ende3e4ddc183ca1a46598b11c2d0d13966:
		;
	case BlockAMD64UGT:
		// match: (UGT (InvertFlags cmp) yes no)
		// cond:
		// result: (ULT cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto end49818853af2e5251175d06c62768cae7
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64ULT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end49818853af2e5251175d06c62768cae7
	end49818853af2e5251175d06c62768cae7:
		;
	case BlockAMD64ULE:
		// match: (ULE (InvertFlags cmp) yes no)
		// cond:
		// result: (UGE cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto endd6698aac0d67261293b558c95ea17b4f
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64UGE
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto endd6698aac0d67261293b558c95ea17b4f
	endd6698aac0d67261293b558c95ea17b4f:
		;
	case BlockAMD64ULT:
		// match: (ULT (InvertFlags cmp) yes no)
		// cond:
		// result: (UGT cmp yes no)
		{
			v := b.Control
			if v.Op != OpAMD64InvertFlags {
				goto end35105dbc9646f02577167e45ae2f2fd2
			}
			cmp := v.Args[0]
			yes := b.Succs[0]
			no := b.Succs[1]
			b.Kind = BlockAMD64UGT
			b.Control = cmp
			b.Succs[0] = yes
			b.Succs[1] = no
			return true
		}
		goto end35105dbc9646f02577167e45ae2f2fd2
	end35105dbc9646f02577167e45ae2f2fd2:
	}
	return false
}