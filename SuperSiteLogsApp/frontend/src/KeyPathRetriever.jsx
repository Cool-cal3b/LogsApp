import React, { useEffect, useState } from "react";
import { GetKeyName, SetKey } from '../wailsjs/go/handlers/KeyPathSelect'

function KeyPathSelect() {
    const [fileName, setFileName] = useState('')

    useEffect(() => {
        async function fetchFileName(params) {
            try {
                let fileNameFromServer = await GetKeyName();
                setFileName(fileNameFromServer)
            } catch (err) {
                console.error('Error fetching data:', err);
            }
        }
        
        fetchFileName();
    }, []);

    async function handleFileChange(event) {
        const file = event.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = async function(e) {
                const content = e.target.result;
                const base64Content = btoa(String.fromCharCode(...new Uint8Array(content)));
                await SetKey(base64Content, file.name); // Pass the base64 string to the backend
                setFileName(file.name);  // Display the file name
            };
            reader.readAsArrayBuffer(file); // Reads the file as binary data
        }
    }

    return(
        <div>
            <div>
                <input id="PathSelect" type="file" style={{display: 'none'}} onChange={handleFileChange} />
                <label htmlFor="PathSelect">Choose file</label>
            </div>

            <div>{fileName}</div>
        </div>
    );
}

export default KeyPathSelect;