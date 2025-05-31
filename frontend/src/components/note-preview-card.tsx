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
        <Card className="binding-border min-w-2xs w-full h-full border transition-colors duration-200 hover:border-primary hover:cursor-pointer">
            <CardHeader>
                <CardTitle>
                    {note.title}
                </CardTitle>
                <CardDescription>
                    {new Date(note.created_at).toLocaleString()}
                </CardDescription>
                <CardContent className="text-limit-5">
                    {note.content}
                </CardContent>
            </CardHeader>
        </Card>
    )
}