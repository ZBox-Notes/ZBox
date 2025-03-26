import { useAllNotes } from '@/api/queries/notes-queries';
import { NotePreviewCard } from '@/components/note-preview-card';
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
            {isSuccess ? (
                (data && data.length > 0) ? (
                    data.map(note => (
                        <NotePreviewCard key={note.id} note={note} />
                    ))
                ) : (
                    <div>
                        <p className='italic'>
                            No notes yet
                        </p>
                    </div>
                )
            ) : (
                <div>
                    <p className='italic'>
                        Oops... Unable to fetch notes
                    </p>
                </div>
            )}
        </div>
    );
};

export default NotesPage;
