import React, {useEffect, useState} from "react";
import {SignForm} from "./signForm";
import {Chat} from "./chat/chat";
import axios from "axios";
import {CircularProgress} from "@material-ui/core";

export function TopPage(props) {
    const [auth, setAuth] = useState(false)
    const [load, setLoad] = useState(true)

    useEffect(() => {
        verify().then((ok) => {
            setAuth(ok)
            setLoad(false)
        }).catch((err) => {
            setAuth(false)
            setLoad(false)
        })
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
            {load ? <CircularProgress color="secondary"/> : (auth ? <Chat/> : <SignForm/>)}
        </div>

    );
}