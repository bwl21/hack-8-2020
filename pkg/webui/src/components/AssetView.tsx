import * as React from 'react';
import { withStyles, WithStyles, createStyles, Theme, Table, TableRow, TableCell, Toolbar, Button } from '@material-ui/core';
import { AssetComponent } from '../generated/graphql';

const styles = (theme: Theme) => createStyles({
    label: {
        fontWeight: "bold"
    }
});

export interface AssetViewProps extends WithStyles {
    filename: string
}

const AssetViewImpl: React.FunctionComponent<AssetViewProps> = (props) => {
    return <AssetComponent variables={{filename: props.filename}}>{result => 
        <React.Fragment>
            { result.loading && <div>loading</div> }
            { result.error && <div>{ result.error.message }</div> }
            { result.data && result.data.asset && <React.Fragment>
                <Toolbar>
                    <Button href={`https://zn-next.zupfnoter.de/?load=${result.data.asset.media.self}`}>Preview</Button>
                </Toolbar>
                <Table>
                    <TableRow>
                        <TableCell className={props.classes.label}>ID</TableCell>
                        <TableCell>{result.data.asset.id}</TableCell>
                    </TableRow>
                </Table>
            </React.Fragment> }
        </React.Fragment>
    }</AssetComponent>;
};

export const AssetView = withStyles(styles)(AssetViewImpl);
