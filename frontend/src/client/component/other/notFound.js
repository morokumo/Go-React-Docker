import React, {useEffect} from "react";
import {Redirect} from "react-router-dom";

export function NotFound(props) {

    useEffect(() => {
    }, [])

    return (
        <Redirect to={'/'}/>
    );
}