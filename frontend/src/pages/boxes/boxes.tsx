import { Archive } from 'lucide-react';
import React from 'react';

const Boxes: React.FC = () => {
    return (
        <div>
            <div className='flex items-center align-center space-x-4'>
                <Archive size={48} />
                <h1>Boxes</h1>
            </div>
            {/* Add your content here */}
        </div>
    );
};

export default Boxes;
