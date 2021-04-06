import React, {useEffect, useState} from "react";
import {ChatRoom} from "./room/chatRoom";
import LeftNav from "./leftNav";

import {io} from "socket.io-client";
import {AccountList} from "./accountList";

export function Chat(props) {
    const [room, setRoom] = useState(undefined)
    const [socket, setSocket] = useState(undefined)
    useEffect(()=>{
        setSocket(io({path: '/api/web/'}))
    },[])
    return (
        <div>
            <table>
                <tr>
                    <td>
                        <LeftNav room={room} setRoom={setRoom}/>
                    </td>
                    <td>
                        {room !== undefined ? <ChatRoom p={props} room={room} socket={socket}/> : ''}
                    </td>
                    <td>
                        {room !== undefined ? <AccountList p={props} room={room} socket={socket}/> : ''}
                    </td>
                </tr>
            </table>
        </div>
    )

}