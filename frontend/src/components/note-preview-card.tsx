import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle
} from "@/components/ui/card";
import { Note } from "@/types/types";
import React from "react";

interface NotePreviewCardProps {
    note: Note
}

export const NotePreviewCard: React.FC<NotePreviewCardProps> = (props: { note: Note }): JSX.Element => {
    let note = props.note
    return (
        <Card>
            <CardHeader>
                <CardTitle>
                    {note.title}
                </CardTitle>
                <CardDescription>
                    {note.created_at}
                </CardDescription>
                <CardContent>
                    {note.content}
                </CardContent>
            </CardHeader>
        </Card>
    )
}