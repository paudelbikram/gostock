// src/utils/numberFormat.js
export const formatSmartNumber = (value) => {
    const abs = Math.abs(value);
    if (abs >= 1e12) return (value / 1e12).toFixed(2) + 'T';
    if (abs >= 1e9) return (value / 1e9).toFixed(2) + 'B';
    if (abs >= 1e6) return (value / 1e6).toFixed(2) + 'M';
    if (abs >= 1e3) return (value / 1e3).toFixed(2) + 'K';
    return Number(value).toLocaleString(); // fallback with commas
  };
  
  export const formatSmartMoney = (value) => {
    return '$' + formatSmartNumber(value);
  };
  