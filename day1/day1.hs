main :: IO ()
main = do
    ints <- readInts "day1.txt"
    printGreater ints
    printGreater . windows $ ints
    where
        printGreater = print . countGreater . pairs

readInts :: String -> IO [Int]
readInts name = do
    file <- readFile name
    return $ readInts' file
    where
        readInts' = map (read :: String -> Int) . lines 

pairs :: [a] -> [(a, a)]
pairs ls = zip ls (tail ls)

windows :: [Int] -> [Int]
windows [] = []
windows ls =
    let headSum = sum $ take 3 ls
        tailWindows = windows $ drop 1 ls
    in headSum : tailWindows

countGreater :: (Ord a) => [(a, a)] -> Int
countGreater = length . filter (uncurry (<))
