import React, {useEffect, useState} from "react";
import {SignForm} from "./signForm";
import {Chat} from "./chat/chat";
import axios from "axios";
import {Redirect} from "react-router-dom";

export function TopPage(props) {
    const [auth, setAuth] = useState(true)

    useEffect(() => {
        verify().then((ok) => {
            setAuth(ok)
        }).catch((ok) => {
            setAuth(false)
        })
        console.log(props)
    }, [props])


    async function verify() {
        let ok = false
        await axios.post('api/auth')
            .then((response) => {
                if (response.data.message === 'Status OK') {
                    ok = true;
                }
            })
        return ok
    }

    return (
        <div>
            {auth ? <Chat/> :  <SignForm/>}
        </div>

    );
}