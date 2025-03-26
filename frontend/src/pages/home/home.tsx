import { useQOTD } from '@/api/queries/quotes-queries';
import { House } from 'lucide-react';
import React from 'react';

const Home: React.FC = () => {
    const { data, error } = useQOTD()

    return (
        <div>
            <div className='flex items-center align-center space-x-4'>
                <House size={48} strokeWidth={1} />
                <h1>Home</h1>
            </div>
            {error ? <></> : <blockquote>
                {data?.a}
            </blockquote>}
        </div>
    );
};

export default Home;
