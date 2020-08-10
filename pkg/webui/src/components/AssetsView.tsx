import * as React from 'react';
import { withStyles, WithStyles, createStyles, Theme, Table, TableRow, TableCell, TableHead, TableBody, Link } from '@material-ui/core';
import { Asset, AssetsComponent } from '../generated/graphql';

const styles = (theme: Theme) => createStyles({
    
});

export interface AssetsViewProps extends WithStyles {
}

interface AssetsViewState {
    search?: string;
}

class AssetsViewImpl extends React.Component<AssetsViewProps, AssetsViewState> {

    constructor(props: AssetsViewProps) {
        super(props);

        this.state = {};
    }

    public render() {
        return <AssetsComponent>
            {result => 
                <div>
                    { result.loading && <div>loading</div> }
                    { result.error && <div>{ result.error.message }</div> }
                    { result.data && <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>ID</TableCell>
                                <TableCell>Title</TableCell>
                                <TableCell>Filename</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            { result.data.assets.map(v => <TableRow key={v.id}>
                                <TableCell><Link href={`/asset/${v.filename}`}>{ v.id }</Link></TableCell>
                                <TableCell>{ v.title }</TableCell>
                                <TableCell>{ v.filename }</TableCell>
                            </TableRow>) }
                        </TableBody>
                    </Table> }
                </div>
            }
        </AssetsComponent>
    }

}

export const AssetsView = withStyles(styles)(AssetsViewImpl);
