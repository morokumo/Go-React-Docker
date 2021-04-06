import React, {useEffect, useState} from "react";
import axios from "axios";
import {Card, CardContent, ListItem, ListItemText, makeStyles} from "@material-ui/core";
import {FixedSizeList} from "react-window";

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        maxWidth: 300,
        minWidth: 100,
        height: 700,
        backgroundColor: theme.palette.background.paper,
    },
    content: {
        width: '100%',
        maxWidth: 780,
        minWidth: 320,
        maxHeight: 600,
        minHeight: 700,
        backgroundColor: theme.palette.background.paper,
    },
    prime: {
        fontSize: 14
    },
    date: {
        fontSize: 12
    },
    second: {
        fontSize: 18
    }
}));

export function AccountList(props) {
    const [load, setLoad] = useState()
    const [room, setRoom] = useState(props.room);
    const [accounts, setAccounts] = useState([]);
    const classes = useStyles()
    useEffect(() => {
        props.socket.on("isOnline", (e) => {

        })
    }, [])

    useEffect(() => {
        getAccounts()
    }, [room]);

    useEffect(() => {
        setRoom(props.room)
    }, [props.room]);


    async function getAccounts() {
        await axios.post('/api/getRoomAccounts', {'room_id': room.ID})
            .then(res => {
                if (!(res.data === null) && !(res.data === undefined)) {
                    setAccounts(res.data.accounts)
                }
            }).catch(error => {

            })
    }

    const Row = ({index, style}) => (
        <div>
            <ListItem style={style}>
                <ListItemText primary={accounts[index].ID}/>
            </ListItem>
        </div>
    );

    return (
        <div>
            <Card className={classes.root}>
                <CardContent>
                    <h2>account list</h2>
                    <FixedSizeList height={400} width={600} itemSize={60} itemCount={accounts.length}>
                        {Row}
                    </FixedSizeList>
                </CardContent>
            </Card>


        </div>
    );
}