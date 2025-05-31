import NotesList from '@/components/notes-list';
import Quote from '@/components/quote';
import React from 'react';

const Home: React.FC = () => {

    return (
        <div>
            <div className='flex items-center align-center space-x-4'>
                <h1>Home</h1>
            </div>
            <div className='flex flex-row w-full'>
                <div className='ml-auto'>
                    <Quote />
                </div>
            </div>

            <div className='flex items-center align-center space-x-4'>
                <h2>My notes</h2>
            </div>
            <NotesList />
        </div>
    );
};

export default Home;
