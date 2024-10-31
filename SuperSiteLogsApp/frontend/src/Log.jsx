import { Link, useParams } from "react-router-dom";
import { GetLog } from '../wailsjs/go/handlers/LogView'
import { useEffect, useState } from "react";
import styles from './pageCss/LogView.module.css'

function Log() {
    const { id } = useParams();
    const [message, setMessage] = useState('');
    const [date, setDate] = useState('');
    const [userName, setUsername] = useState('');

    useEffect(() => {
        async function getLogInfo() {
            let logInfo = await GetLog(parseInt(id));
            setMessage(logInfo.Log.Message)
            setDate(formatDate(logInfo.Log.Time))
            setUsername(logInfo.UserName)
        }
        getLogInfo()
    }, [id])

    function formatDate(date) {
        return new Date(date).toLocaleString('en-US', {
            month: 'short',
            day: 'numeric',
            year: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            hour12: true
          })
    }

    function isHtml(content) {
        const div = document.createElement('div');
        div.innerHTML = content;
        return div.innerHTML !== content;
    }

    return (
        <div className={styles.outterCont}>
            <div className={ styles.cont }>
                <Link to={'/Logs'} className={styles.backButton}>Back</Link>
                <div className={ styles.topText}>{ userName }</div>
                <div className={ styles.topText}>{ date }</div>
                <div className={ styles.messageBox}>
                    {
                        isHtml(message) ? (
                            <div dangerouslySetInnerHTML={{ __html: message }}></div>
                        ) : (
                            <div>{message}</div>
                        )
                    }
                </div>
            </div>
        </div>
    )
}

export default Log;