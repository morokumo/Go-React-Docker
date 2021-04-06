import React, {useEffect, useState} from "react";
import axios from "axios";
import {ListItem, ListItemIcon, ListItemText, makeStyles} from "@material-ui/core";
import KeyboardArrowRightIcon from '@material-ui/icons/KeyboardArrowRight';
import {FixedSizeList} from 'react-window';
import {VisibilityOffOutlined} from "@material-ui/icons";

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        maxWidth: 360,
        minWidth: 300,
        backgroundColor: theme.palette.background.paper,
    },
}));


export function FindMyRooms(props) {
    const [rooms, setRooms] = useState([])
    const classes = useStyles()
    useEffect(() => {
        findMyRooms()
    }, []);


    function findMyRooms() {
        axios.post(`/api/findMyRooms`)
            .then(res => {
                if (!(res.data.rooms === null) && !(res.data.rooms === undefined)) {
                    setRooms(res.data.rooms)
                }
            })
    }

    const Row = ({index, style}) => (

        <ListItem style={style} button onClick={() => props.p.setRoom(rooms[index])}>
            <ListItemText
                primary={
                    rooms[index].Name.length > 10 ?
                        rooms[index].Name.substr(0, 10) + '...' : rooms[index].Name}/>
            {rooms[index].Private ? <ListItemIcon><VisibilityOffOutlined/></ListItemIcon> : ''}
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