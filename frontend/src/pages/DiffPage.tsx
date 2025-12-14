import  { useEffect, useState } from "react";
import { Table, Container } from "semantic-ui-react";

export const DiffPage: React.FC = () => {
    const [diffs, setDiffs] = useState<any[]>([]);

    useEffect(() => {
        fetch("/diff")
            .then(res => res.json())
            .then(setDiffs);
    }, []);

    return (
        <Container>
            <h2>Diff</h2>
            <Table celled>
                <Table.Header>
                    <Table.Row>
                        <Table.HeaderCell>Business Key</Table.HeaderCell>
                        <Table.HeaderCell>Type</Table.HeaderCell>
                        <Table.HeaderCell>Changes</Table.HeaderCell>
                    </Table.Row>
                </Table.Header>
                <Table.Body>
                    {diffs.map((d) => (
                        <Table.Row key={d.businessKey}>
                            <Table.Cell>{d.businessKey}</Table.Cell>
                            <Table.Cell>{d.type}</Table.Cell>
                            <Table.Cell>
                                {d.fields?.map((f:any) => `${f.field}: ${f.old} → ${f.new}`).join(", ")}
                            </Table.Cell>
                        </Table.Row>
                    ))}
                </Table.Body>
            </Table>
        </Container>
    );
};
