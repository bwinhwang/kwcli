#!/bin/bash

# Function to execute a kwcli command and check result
execute_and_check() {
    command="$1"
    expected_output="$2" # Optional

    result=$($command) 

    if [[ $? -ne 0 ]]; then  # Check exit code
        echo "Command failed: $command"
        echo "Result: $result"
        exit 1
    fi

    # Optional: Check output if an expected pattern is provided
    if [[ -n "$expected_output" && ! $result =~ $expected_output ]]; then
        echo "Unexpected output from command: $command"
        echo "Result: $result"
        exit 1
    fi
}

# Your test sequence
execute_and_check "../bin/linux-amd64/kwcli delete_project -n kwcli_test" 
sleep 1
execute_and_check "../bin/linux-amd64/kwcli create_project -n kwcli_test" 
sleep 1
execute_and_check "../bin/linux-amd64/kwcli create_module -p kwcli_test -n module_test --paths \"c\",\"**/a/b/**\" " 
sleep 1 

execute_and_check "../bin/linux-amd64/kwcli create_view -p kwcli_test -n view_test --query \"code:npd,abv,iter.,misra.goto.,nnts,naming.3,CL.SELF-ASSIGN,CWARN.FUNCADDR,FMM,FNH.,funcret,locret,MISRA.ADDR.REF.PARAM,MISRA.CATCH.,MISRA.DEFINE.NOPARS,MISRA.LITERAL.NULL.INT,MISRA.SWITCH.LABEL,mlk,retvoid,rh.leak,ufm,uninit,voidret,CWARN.DTOR.NONVIRT.,MISRA.COPYASSIGN.ABSTRACT,MISRA.THROW.,UNREACH.RETURN,strbo,conc.,constcond,semicol,effect,CWARN.MEMSET.SIZEOF.PTR,MISRA.FLOAT.BIT.REPR,MISRA.PRAGMA.ASM,FUM.GEN.,sv.fmt,SV.INCORRECT_RESOURCE_HANDLING,INCORRECT.ALLOC_SIZE,passbyvalue,MISRA.COPY.CSTR.TMPL,MISRA.SIZEOF.SIDE_EFFECT,technisat.calling.7,technisat.calling.8,CWARN.OVERRIDE.CONST,MISRA.ASSIGN.COND,MISRA.NULL.STMT,MISRA.FOR.COND.EQ,cwarn.mem.nonpod,TECHNISAT.CALLING.1,TECHNISAT.CALLING.3,MISRA.SWITCH.NODEFAULT,UNREACH.GEN,VA_UNUSED.GEN,LV_UNUSED.GEN,VA_UNUSED.INIT,MISRA.SHIFT.RANGE,MISRA.FUNC.ADDR,RN.INDEX,CWARN.COPY.NOASSIGN,MISRA.ASSIGN.SUBEXPR,MISRA.SWITCH.NO_BREAK,MISRA.LOGIC.OPERATOR.NOT_BOOL,CWARN.NULLCHECK.FUNCNAME,MISRA.FLOAT_EQUAL,tainted,CWARN.HIDDEN.PARAM,MISRA.SAME.DEFPARAMS,MISRA.FUNC.STATIC.REDECL,CWARN.BITOP.SIZE,MISRA.COPYASSIGN.TMPL,MISRA.INCL.SYMS,MISRA.DECL.EXCPT.SPEC,MISRA.CAST.INT_FLOAT,MISRA.EXPANSION.DIRECTIVE,MISRA.FOR.LOOP_CONTROL.CHANGE.COND,MISRA.FOR.LOOP_CONTROL.CHANGE.EXPR,MISRA.TOKEN.BADCOM,MISRA.TOKEN.WRONGESC,MISRA.NS.USING.HEADER,MISRA.NS.USING_DECL,CL.FFM.ASSIGN,CL.FFM.COPY,MISRA.PUREVIRT.OVRD,MISRA.FLOAT_EQUAL,MISRA.BASE.MANYDEFS,MISRA.CTOR.DYNAMIC,MISRA.DTOR.DYNAMIC,MISRA.COPY.CSTR.TMPL,MISRA.SPEC.ILL,MISRA.COPYASSIGN.TMPL,MISRA.FUNC.SPEC.OVRLD,MISRA.SPEC.SAMEFILE,MISRA.TEMPLMEM.NOQUAL,MISRA.NAMESPACE.,MISRA.UMINUS.UNSIGNED,MISRA.CTOR.NOT_EXPLICIT,MISRA.COMMA,INCONSISTENT.LABEL,UNREACH.SIZEOF,NUM.OVERFLOW,INVARIANT_CONDITION.GEN,INVARIANT_CONDITION.UNREACH,CL.MLK.ASSIGN,CL.SHALLOW.COPY,CL.SHALLOW.ASSIGN,DBZ.,SV.TAINTED.SECURITY_DECISION,HCC,RCA,MISRA.FILE_PTR.DEREF.,MISRA.NS.GLOBAL.USING,SV.TAINTED.CALL.DEREF,SV.TAINTED.DEREF,MISRA.FUNC.RECUR,INFINITE_LOOP.LOCAL,CWARN.MEMBER.INIT.ORDER -code:MISRA.CATCH.ALL,technisat.calling.10,misra.iter,CWARN.CONSTCOND.WHILE,PORTING.CMPSPEC.EFFECTS.ASSIGNMENT severity:critical,error -status:Filter,Not a Problem -module:external_code,generated_code,test_code\" " 
sleep 1 

execute_and_check "../bin/linux-amd64/kwcli create_module -p kwcli_test -n module_test --paths \"c\",\"**/a/b/**\" " 
sleep 1 
execute_and_check " ../bin/linux-amd64/kwcli search -p DEV__RTCU-B3_jn-src-service-jn-lifecycle-service --summary" 
sleep 1 



echo "All tests completed successfully!"
