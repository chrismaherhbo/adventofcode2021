import Data.List (foldl')
import Text.Read (readMaybe)

data Direction = Forward | Up | Down deriving (Show, Eq)

data Movement = Movement { direction :: Direction, magnitude :: Int } deriving (Show, Eq)

main :: IO ()
main = readMovements >>= print . calculatePosition

directionFrom :: String -> Maybe Direction
directionFrom "forward" = Just Forward
directionFrom "up" = Just Up
directionFrom "down" = Just Down
directionFrom _ = Nothing

parseMovement :: [String] -> Maybe Movement
parseMovement (x:y:_) = do
    d <- directionFrom x
    m <- readMaybe y :: Maybe Int
    pure $ Movement d m

readMovements :: IO [Maybe Movement]
readMovements = readFile "day2.txt" >>=
    return . map parseMovement . map words . lines

calculatePosition :: [Maybe Movement] -> Int
calculatePosition = calculate . position
    where
        position = foldl' (\(x, y, z) move -> case move of
                Just (Movement Forward m) -> (x + m, y + z * m, z)
                Just (Movement Up m) -> (x, y, z - m)
                Just (Movement Down m) -> (x, y, z + m)
            ) (0, 0, 0)
        calculate (x, y, _) = x * y
