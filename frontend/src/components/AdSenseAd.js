import React, { useEffect, useRef } from 'react';

const AdSenseAd = ({
  adClient = "ca-pub-4388938725058084",
  adSlot = "REPLACE_WITH_SLOT",
  style = { display: "block" },
  format = "auto",
  responsive = true,
  adTest = false
}) => {
  const adRef = useRef(null);

  useEffect(() => {
    if (!window || !adRef.current) return;

    try {
      (window.adsbygoogle = window.adsbygoogle || []).push({});
    } catch (e) {
      console.warn("Ads error:", e);
    }
  }, []);

  return (

    <div style={{ width: "100%" }}>
      {/* SAFE, POLICY-COMPLIANT HEADS UP MESSAGE */}
      <p style={{ fontSize: "14px", color: "#555", marginBottom: "4px" }}>
      You may see ads below to support this website.
      </p>
      <ins
        className="adsbygoogle"
        ref={adRef}
        style={style}
        data-ad-client={adClient}
        data-ad-slot={adSlot}
        data-ad-format={format}
        {...(responsive ? { "data-full-width-responsive": "true" } : {})}
        {...(adTest ? { "data-adtest": "on" } : {})}
      />
    </div>
  );
};

export default AdSenseAd;
