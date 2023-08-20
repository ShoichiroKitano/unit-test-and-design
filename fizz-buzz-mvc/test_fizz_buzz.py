import unittest
import subprocess

class FizzBuzzTest(unittest.TestCase):
    def __program(self):
        p = subprocess.Popen(
                ['go', 'run', 'main.go'],
                stdin=subprocess.PIPE,
                stdout=subprocess.PIPE
                )
        return p

    def test_3(self):
        p = self.__program()
        input = "3"
        result = p.communicate(input.encode())[0].decode()
        self.assertEqual(result, 'FAizz\n')
