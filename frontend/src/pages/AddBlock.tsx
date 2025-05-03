import React, { useState } from 'react';

const AddBlockPage: React.FC = () => {
  const [data, setData] = useState('');

  const handleSubmit = () => {
    alert(`Block added with data: ${data}`);
  };

  return (
    <div className="p-4">
      <h2 className="text-2xl font-semibold mb-4">Add New Block</h2>
      <input
        type="text"
        placeholder="Block Data"
        className="border p-2 mr-2"
        value={data}
        onChange={(e) => setData(e.target.value)}
      />
      <button onClick={handleSubmit} className="bg-blue-500 text-white px-4 py-2 rounded">
        Add Block
      </button>
    </div>
  );
};

export default AddBlockPage;
