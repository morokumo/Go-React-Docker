import React, {useEffect, useState} from "react";
import axios from "axios";
import {Button, ListItem, ListItemIcon, ListItemText} from "@material-ui/core";
import KeyboardArrowRightIcon from "@material-ui/icons/KeyboardArrowRight";
import {FixedSizeList} from "react-window";


export function FindPublicRooms(props) {
    const [rooms, setRooms] = useState([])
    useEffect(() => {
        findRooms()
    }, []);

    function findRooms() {
        axios.post(`/api/findPublicRooms`)
            .then(res => {
                console.log(res.data)
                if (!(res.data.rooms === null) && !(res.data.rooms === undefined)) {
                    setRooms(res.data.rooms)
                }
            })

    }

    async function joinRoom(event) {
        let result = confirm("join this room ?")
        if (result) {
            let data = {"room_id": event.target.value}
            console.log("DATA:", event)
            await axios.post(`/api/joinRoom`, data)
                .then(res => {

                })
        }

        findRooms()
    }

    const Row = ({index, style}) => (
        <ListItem style={style}>
            <ListItemText
                primary={rooms[index].Name.length > 10 ? rooms[index].Name.substr(0, 10) + '...' : rooms[index].Name}/>
            <button color="secondary" value={rooms[index].ID} onClick={joinRoom}>JOIN</button>
            <ListItemIcon>
                <KeyboardArrowRightIcon/>
            </ListItemIcon>
        </ListItem>
    );
    return (
        <div>
            <FixedSizeList height={400} width={300} itemSize={46} itemCount={rooms.length}>
                {Row}
            </FixedSizeList>
        </div>

    )

}