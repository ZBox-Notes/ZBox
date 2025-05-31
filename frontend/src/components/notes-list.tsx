import { useAllNotes } from '@/api/queries/notes-queries';
import { Note } from '@/types/types';
import React from 'react';
import { NotePreviewCard } from './note-preview-card';

const NotesList: React.FC = () => {
    let notes: Note[] = [];

    const { data, isLoading, isError } = useAllNotes();

    if (!isError) {
        notes = data || [];
    }

    if (isLoading) {
        return <></>
    }

    return (
        <div className='flex flex-row gap-4 flex-wrap'>
            {isLoading ? <>Loading</> : notes.map((note) => (
                <div className='flex-1' key={note.id}>
                    <NotePreviewCard note={note} />
                </div>
            ))}
        </div>
    );
};

export default NotesList;