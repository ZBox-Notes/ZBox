import { useAllNotes } from '@/api/queries/notes-queries';
import NotesList from '@/components/notes-list';
import { StickyNote } from 'lucide-react';
import React from 'react';

const NotesPage: React.FC = () => {

    const { data, isSuccess, error } = useAllNotes()
    if (error) console.log(error);

    return (
        <div>
            <div className='flex items-center align-center space-x-4 mb-4'>
                <StickyNote size={48} strokeWidth={1} />
                <h1>Notes</h1>
            </div>
            <NotesList />
        </div>
    );
};

export default NotesPage;
