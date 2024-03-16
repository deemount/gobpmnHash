/*
 * MIT License
 *
 * Copyright (c) 2024 Salvatore Gonda
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*

Module gobpmnHash is a simple implementation of the BPMN 2.0 hash algorithm.

Hash values are used to generate unique IDs for each element of a process.

	* The hash value is generated using the crypto/rand package and the hash/fnv package to generate a 32-bit FNV-1a hash.
	* The suffix is used to generate a unique ID for each element of a process.

Injections are used to inject hash values into the fields of a struct.

	* The injectConfig method sets the bool type.
	* The injectCurrentField method injects the current field with a hash value.
	* The injectNextField method injects the next field with a hash value.

*/

package gobpmn_hash
