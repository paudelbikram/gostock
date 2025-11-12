import React, { useState } from "react";
import { useSelector } from 'react-redux';

import {
  Box,
  TextField,
  Typography,
  Button,
  Card,
  CardContent,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  FormHelperText
} from "@mui/material";
import { formatSmartNumber } from "../utils/numberFormat";

const DCFCalculator = () => {
  const { data } = useSelector((state) => state.stock);
  const { cashflow, overview } = data;
  const annualReportsCashflow = cashflow.annualReports;
  const initialFCF = Number(annualReportsCashflow[0]?.operatingCashflow || 0) - Math.abs(Number(annualReportsCashflow[0]?.capitalExpenditures || 0));;
  const sharesOutstanding = overview.SharesOutstanding;
  const terminalGrowth = 3;
  const discountRate = 10;
  const [inputs, setInputs] = useState({
    growthRate: 10,
    years: 10,
  });


  const [result, setResult] = useState({
    fairValuePerShare: 0,
    totalPV: 0,
  });

  const handleSubmit = () => {
    const growthRate = parseFloat(inputs.growthRate) || 0;
    const years = parseInt(inputs.years) || 0;
    const g = growthRate / 100;
    const r = discountRate / 100;
    const tg = terminalGrowth / 100;

    let totalPV = 0;
    let fcf = initialFCF;

    for (let i = 1; i <= years; i++) {
      fcf *= 1 + g; // project FCF
      totalPV += fcf / Math.pow(1 + r, i);
    }

    const terminalValue = (fcf * (1 + tg)) / (r - tg);
    totalPV += terminalValue / Math.pow(1 + r, years);

    const fairValuePerShare = sharesOutstanding
      ? totalPV / sharesOutstanding
      : 0;
    setResult({ fairValuePerShare, totalPV });
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setInputs((prev) => ({ ...prev, [name]: value }));
  };
  
  return (
    <Card sx={{ maxWidth: 500, mx: "auto", mt: 4, p: 2, boxShadow: 3 }}>
      <CardContent>
        <Typography variant="h5" gutterBottom align="center">
          Stock Price Calculator
        </Typography>
        <Box display="flex" flexDirection="column" gap={2}>
          <TextField
            label="Growth Rate (%)"
            name="growthRate"
            value={inputs.growthRate}
            onChange={handleChange}
            type="number"
            helperText="Expected annual FCF growth, e.g. 10 for 10%"
          />
          <FormControl fullWidth variant="outlined">
            <InputLabel id="projection-years-label">Projection Years</InputLabel>
            <Select
              labelId="projection-years-label"
              name="years"
              value={inputs.years}
              onChange={handleChange}
              label="Projection Years"
            >
              {[5, 10, 15, 20].map((year) => (
                <MenuItem key={year} value={year}>
                  {year} Years
                </MenuItem>
              ))}
            </Select>
            <FormHelperText>Select number of years for cash flow projection</FormHelperText>
          </FormControl>
          <FormHelperText>Assumes 10% discount rate, 3% terminal growth, initial FCF of {formatSmartNumber(initialFCF)} {overview.Currency}, and {formatSmartNumber(sharesOutstanding.toLocaleString())} shares outstanding.</FormHelperText>
          <Button variant="contained" color="primary" onClick={handleSubmit} >
            Calculate
          </Button>
        </Box>

        <Box mt={3} textAlign="center">
          <Typography variant="h6">
            Estimated Fair Value per Share:{" "}
            <strong>{result.fairValuePerShare.toFixed(2)} {overview.Currency}</strong>
          </Typography>
          <Typography variant="body2" color="textSecondary">
            Total Present Value of Cash Flows: $
            {formatSmartNumber(result.totalPV)}
          </Typography>
        </Box>
      </CardContent>
    </Card>
  );
};

export default DCFCalculator;
