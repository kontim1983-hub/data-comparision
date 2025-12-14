import React from "react";
import {DiffPage} from "./pages/DiffPage";
import {ImportPage} from "./pages/ImportPage";

const App: React.FC = () => {
    return (
        <div>
            <h1>Leasing Data Comparator</h1>
            <DiffPage/>
            <ImportPage/>
        </div>
    );
};

export default App;
