"use client";

import React, { createContext, useContext, useState, useEffect } from "react";

const ThemeContext = createContext();

export function ThemeProvider({ children }: { children: React.ReactNode }) {
  const [mode, setMode] = useState('');
}
