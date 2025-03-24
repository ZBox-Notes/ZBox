import { StickyNote } from 'lucide-react';
import React from 'react';

const Notes: React.FC = () => {
    return (
        <div>
            <div className='flex items-center align-center space-x-4'>
                <StickyNote size={48} />
                <h1>Notes</h1>
            </div>
            {/* Add your content here */}
        </div>
    );
};

export default Notes;
