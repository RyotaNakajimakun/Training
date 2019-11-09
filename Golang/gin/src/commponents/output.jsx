import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import CardContent from '@material-ui/core/CardContent'
import Card from '@material-ui/core/Card'
import Container from "@material-ui/core/Container";


const useStyles = makeStyles({
    card: {
        minWidth: 275,
    },
    bullet: {
        display: 'inline-block',
        margin: '0 2px',
        transform: 'scale(0.8)',
    },
    title: {
        fontSize: 14,
    },
    pos: {
        marginBottom: 12,
    },
});


export default class RenderModules extends React.Component {
    render() {
        const classes = useStyles;
        const modules = this.props.values;

        return (
            <Container maxWidth="sm">
                {modules.map((module, i) => (
                    <Card className={classes.card} style={{marginTop: 10}}>
                        <CardContent>
                            {module.value.map((value, key) => (
                                <Typography variant="body1" component="p" key={key}>
                                    {value}
                                </Typography>
                            ))}
                        </CardContent>
                    </Card>
                ))}
            </Container>
        );
    }
}
