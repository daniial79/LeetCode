def can_construct(ransomNote: str, magazine: str) -> bool:
    alphabetFrequencyTracker = [0] * 26
    
    for i in magazine:
        alphabetFrequencyTracker[ord(i) - 65] += 1
        
    for j in ransomNote:
        alphabetFrequencyTracker[ord(j) - 65] -= 1
        
    for k in alphabetFrequencyTracker:
        if k < 0: return False
        
    return True