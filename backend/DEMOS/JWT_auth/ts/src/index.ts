import express, { Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";
import bcrypt from "bcryptjs";

// ==========================================
// 1. Configuration & Types
// ==========================================

const PORT = 3000;
const JWT_SECRET = "super-secret-key-change-me-in-production";

interface User {
  username: string;
  passwordHash: string;
}

interface AuthJwtPayload {
  username: string;
}

// Extend Express Request interface to include authenticated user information
declare global {
  namespace Express {
    interface Request {
      user?: AuthJwtPayload;
    }
  }
}

// ==========================================
// 2. In-Memory User Store (Demo Purpose)
// ==========================================

const users = new Map<string, User>();

// ==========================================
// 3. JWT Helpers (Sign & Verify)
// ==========================================

// generateToken creates a signed JWT string valid for 1 hour
function generateToken(username: string): string {
  return jwt.sign({ username }, JWT_SECRET, {
    expiresIn: "1h",
  });
}

// ==========================================
// 4. Authentication Middleware
// ==========================================

// authMiddleware checks for the Authorization: Bearer <token> header,
// verifies the JWT, and attaches decoded user to req.user
function authMiddleware(req: Request, res: Response, next: NextFunction): void {
  const authHeader = req.headers.authorization;

  if (!authHeader) {
    res.status(401).json({ error: "Authorization header missing" });
    return;
  }

  const [scheme, token] = authHeader.split(" ");
  if (scheme !== "Bearer" || !token) {
    res.status(401).json({
      error: "Invalid Authorization header format. Expected 'Bearer <token>'",
    });
    return;
  }

  try {
    const decoded = jwt.verify(token, JWT_SECRET) as AuthJwtPayload;
    // Attach decoded user to the request object for downstream route handlers
    req.user = decoded;
    next();
  } catch (error) {
    res.status(401).json({ error: "Invalid or expired token" });
  }
}

// ==========================================
// 5. Express App & Routes
// ==========================================

const app = express();
app.use(express.json());

// Public Route: Sign Up
app.post("/signup", async (req: Request, res: Response): Promise<void> => {
  const { username, password } = req.body;

  if (!username || !password || typeof username !== "string" || typeof password !== "string") {
    res.status(400).json({ error: "Username and password are required" });
    return;
  }

  if (users.has(username)) {
    res.status(409).json({ error: "User already exists" });
    return;
  }

  // 1. Hash password with bcrypt before saving
  const passwordHash = await bcrypt.hash(password, 10);
  users.set(username, { username, passwordHash });

  // 2. Issue a JWT token immediately upon signup
  const token = generateToken(username);

  res.status(201).json({
    message: "User registered successfully",
    token,
    user: { username },
  });
});

// Public Route: Sign In
app.post("/signin", async (req: Request, res: Response): Promise<void> => {
  const { username, password } = req.body;

  if (!username || !password) {
    res.status(400).json({ error: "Username and password are required" });
    return;
  }

  const user = users.get(username);
  if (!user) {
    res.status(401).json({ error: "Invalid username or password" });
    return;
  }

  // 1. Compare candidate password with stored bcrypt hash
  const isMatch = await bcrypt.compare(password, user.passwordHash);
  if (!isMatch) {
    res.status(401).json({ error: "Invalid username or password" });
    return;
  }

  // 2. Issue JWT token
  const token = generateToken(username);

  res.status(200).json({
    message: "Login successful",
    token,
    user: { username },
  });
});

// Protected Route: Profile (Requires valid JWT)
app.get("/profile", authMiddleware, (req: Request, res: Response): void => {
  res.status(200).json({
    message: "Welcome to your protected profile!",
    user: req.user,
    timestamp: new Date().toISOString(),
  });
});

// ==========================================
// 6. Server Initialization
// ==========================================

app.listen(PORT, () => {
  console.log(`TypeScript JWT Server running on http://localhost:${PORT}`);
});
