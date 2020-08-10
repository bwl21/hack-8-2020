import * as React from 'react';
import { withStyles, WithStyles, createStyles, Theme, AppBar, Toolbar, Grid, Button } from '@material-ui/core';

const styles = (theme: Theme) => createStyles({
    
});

export interface ApplicationFrameProps extends WithStyles { }

const ApplicationFrameImpl: React.FunctionComponent<ApplicationFrameProps> = (props) => {
    return <React.Fragment>
        <AppBar
            color="primary"
            position="sticky"
        >
            <Toolbar>
                <Grid container>
                    <Grid item xs={2}>
                        Zupfmanager
                    </Grid>
                    <Grid item>
                        <Button href="/assets">Assets</Button>
                        <Button href="/projects">Projects</Button>
                    </Grid>
                </Grid>    
            </Toolbar>    
        </AppBar>
        {props.children}
    </React.Fragment>
}

export const ApplicationFrame = withStyles(styles)(ApplicationFrameImpl);
