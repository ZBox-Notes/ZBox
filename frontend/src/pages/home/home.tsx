import { House } from 'lucide-react';
import React from 'react';

const Home: React.FC = () => {
    return (
        <div>
            <div className='flex items-center align-center space-x-4'>
                <House size={48} />
                <h1>Home</h1>
            </div>
            {/* Add your content here */}
        </div>
    );
};

export default Home;
