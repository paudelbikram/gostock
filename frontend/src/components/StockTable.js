import React from 'react';
import {
  Box,
  Typography,
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Paper,
} from '@mui/material';

import { formatSmartNumber } from '../utils/numberFormat';

const StockTable = ({columns, rows, label}) => {
  return (
    <Box p={2}>
      {/* Responsive Table */}
      <Box mt={4}>
        <Typography variant="h6">{label}</Typography>
        <TableContainer component={Paper} sx={{ mt: 1 }}>
          <Table size="small">
            <TableHead>
              <TableRow>
                {columns.map((col) => (
                <TableCell key={col.id} align={col.align || 'left'}>
                    <strong>{col.label}</strong>
                </TableCell>
                ))}
              </TableRow>
            </TableHead>
            <TableBody>
                {rows.map((row, idx) => (
                <TableRow key={idx}>
                    {columns.map((col) => (
                        <TableCell key={col.id} align={col.align || 'left'}>
                         {col.id === 'fiscalDateEnding' ? row[col.id] : formatSmartNumber(row[col.id])}
                        </TableCell>
                    ))}
                </TableRow>
                ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Box>
    </Box>
  );
};
export default StockTable;
