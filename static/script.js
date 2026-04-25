async function setupTimer(apiPath, labelId, clockId) {
    try {
        const response = await fetch(apiPath);
        if (!response.ok) throw new Error('Rate limit or Server error');
        const data = await response.json();
        document.getElementById(labelId).innerText = data.label;
        const startDate = new Date(data.start_date);

        function updateClock() {
            const now = new Date();
            const diff = now - startDate;

            if (diff < 0) {
                document.getElementById(clockId).innerHTML = "<span style='color: #ef4444'>Target date not reached</span>";
                return;
            }

            const days = Math.floor(diff / (1000 * 60 * 60 * 24));
            const hours = Math.floor((diff / (1000 * 60 * 60)) % 24);
            const minutes = Math.floor((diff / (1000 * 60)) % 60);
            const seconds = Math.floor((diff / 1000) % 60);

            document.getElementById(clockId).innerHTML = `
                <div class="time-segment"><span class="time-value">${days}</span><span class="unit">Days</span></div>
                <div class="separator">:</div>
                <div class="time-segment"><span class="time-value">${hours.toString().padStart(2, '0')}</span><span class="unit">Hours</span></div>
                <div class="separator">:</div>
                <div class="time-segment"><span class="time-value">${minutes.toString().padStart(2, '0')}</span><span class="unit">Mins</span></div>
                <div class="separator">:</div>
                <div class="time-segment"><span class="time-value">${seconds.toString().padStart(2, '0')}</span><span class="unit">Secs</span></div>
            `;
        }

        updateClock();
        setInterval(updateClock, 1000);
    } catch (error) {
        document.getElementById(labelId).innerText = "Connection Error";
        document.getElementById(clockId).innerHTML = `<span style='color: #ef4444; font-size: 0.9rem;'>${error.message === 'Rate limit or Server error' ? 'Too many requests. Please try again later.' : 'Could not connect to the server.'}</span>`;
    }
}

// Khởi tạo các đồng hồ
document.addEventListener('DOMContentLoaded', () => {
    setupTimer('/api/timer1', 'label1', 'clock1');
    setupTimer('/api/timer2', 'label2', 'clock2');
});
