# AI 101: Prompt basics (part 2)

![AI 101 course logo](../logo.png)

## Introduction

In the previous guide, you prepared your environment for working with AI models using both local and cloud deployments. In this section, you focus on the basics of prompting. You will learn how to communicate effectively with AI models to obtain useful and structured responses.

Prompting is a practical skill that allows you to guide AI behavior by providing clear instructions, context, and examples.

Understanding this matters because it shows how AI interprets language and meaning. It helps you build practical skills without needing to write code. You will also develop intuition for how large language models organize and structure knowledge. These concepts provide the foundation for working with more advanced AI workflows later on.

## Prompting Theory

Before diving into the practical exercises, it helps to understand the underlying mechanics of effective communication. The transformation of technical documentation from a peripheral support artifact into a core engineering component emphasizes the importance of precision and cognitive alignment. When you prompt an AI, you are applying these same structural methodologies to guide a machine's output.
The Cognitive Pillars (Prompt Specificity & Instruction)

When you write prompts, you are essentially authoring technical instructions for the AI model. The foundational quality of this communication is anchored in the "Five C's" framework: Clarity, Conciseness, Cohesiveness, Completeness, and Correctness.

* Clarity: Clarity ensures that focus remains on solving the technical problem rather than deciphering ambiguous syntax.

* Conciseness: Conciseness serves as a filter for linguistic noise, ensuring that every word performs a functional role.

* Cohesiveness: Cohesiveness ensures that the logical progression of your prompt mirrors the system or concept being described.

* Completeness: Completeness prevents information gaps that can lead to dead-ends in workflows.

* Correctness: Correctness is the bedrock of technical trust and verification.

### Audience Analysis (Role Prompting)

Assigning a role to an AI directly mirrors the practice of audience analysis in technical writing. A critical failure in communication often stems from the "Curse of Knowledge," a cognitive bias where an expert author cannot simulate the perspective of a newcomer.

* Using audience "personas" has become a standard practice to humanize these categories and ensure content resonates with specific constraints.

* By defining a role, you dictate whether the AI should behave like an expert who values high-information density or a non-specialist who requires significant conceptual grounding.

* Audience analysis is the most critical phase of planning, as it determines the depth of detail, vocabulary choice, and structural complexity of the content.

### Structural Discipline (Multi-Step & Iteration)

Complex prompting requires breaking tasks down into structured workflows. This reflects the adoption of advanced information architectures, such as the Diátaxis framework, which organizes content by user need rather than product feature.

* This systematic approach assigns a single intent to each piece of information, preventing the "muddling" of content types that causes confusion.

* Historical data indicates that "poor organization" is consistently rated as the top issue in engineering writing, often resulting from a failure to create a detailed outline.

* By structuring your prompts and using iterative roadmaps, you avoid disorganized narratives and ensure a clear, logical output.

### Machine Readability (Output Formatting)

Understanding how AI processes text helps you format your prompts for optimal results. A transformative shift in technical communication is the emergence of AI agents as a primary audience for documentation.

* Providing structured, machine-readable data simplifies the ingestion process for the "token-based brains" of AI models.

* As the industry moves toward "answer-first" experiences, utilizing generative engine optimization principles ensures greater control over the output.

* This structured approach reduces the risk of AI "hallucinations" and increases the accuracy of the generated responses.

## Practical exercises

### Understanding prompt specificity

**Task**

Give the AI two prompts for the same topic: one vague and one specific.

> **Example:**
> 
> Vague prompt:
> ```text
> Explain photosynthesis.
> ```
> 
> Specific prompt:
> ```text
> Explain photosynthesis to a 12-year-old using a cooking analogy.
> ```

Compare how the outputs differ in clarity and structure.

### Role prompting

Assign a role to the AI to guide the perspective of the answer.

> **Examples:**
> - `Act as a physics teacher and explain gravity.`
> - `Act as a UX designer reviewing a mobile app interface.`
> - `Act as a cybersecurity expert explaining phishing attacks.`

Observe how the role changes the style and depth of the response.

### Instruction prompting

Rewrite a vague request into a structured instruction.

Example structure:
- Goal
- Constraints
- Output format

> **Example prompt:**
> ```text
> Explain how cloud computing works. Use bullet points and include a real-world example.
> ```

### Few-shot prompting

Provide examples so the model can learn the desired pattern.

> **Example:**
> 
> Input: `Good morning` → Output: `Formal greeting`
> Input: `Thanks a lot` → Output: `Expression of gratitude`
> 
> **Prompt:**
> ```text
> Classify the phrase: 'See you later'.
> ```

### Chain-of-thought prompting

Ask the AI to reason step by step when solving a problem.

> **Example comparison:**
> 
> Prompt 1:
> ```text
> Solve this problem.
> ```
> 
> Prompt 2:
> ```text
> Solve this problem and explain your reasoning step by step.
> ```

Compare the quality of the answers.

### Prompt iteration

Improve a prompt through multiple revisions.

> **Example progression:**
> 1. `Write about renewable energy.`
> 2. `Explain renewable energy sources and their benefits.`
> 3. `Explain three renewable energy sources, their advantages, and real-world examples using bullet points.`

### Multi-step prompting

Break complex tasks into multiple steps.

Example workflow:
1. Ask the AI to create an outline.
2. Expand the outline into sections.
3. Generate the final text.

> **Example prompt:**
> ```text
> Create an outline for an article about artificial intelligence in healthcare and then write a short article based on that outline.
> ```

### Output evaluation

After generating responses, evaluate them based on:
- Accuracy
- Clarity
- Structure
- Creativity

### Prompt comparison

Write two different prompts for the same task and compare the results.

> **Example task:**
> `Summarize an article.`
> 
> Prompt A:
> ```text
> Summarize this text.
> ```
> 
> Prompt B:
> ```text
> Summarize this text in five bullet points and highlight the main conclusion.
> ```

## Building a personal prompt library

Create a collection of prompts you can reuse. Organize them into categories such as:
- Writing
- Learning
- Productivity
- Creativity
- Data analysis

## Summary

In this section, you practiced several techniques for interacting with AI models effectively. These techniques include prompt specificity, role prompting, structured instructions, few-shot examples, and multi-step workflows. These exercises help build a foundation for using AI tools in real-world scenarios.

---

Would you like me to help you draft the next part of this AI 101 course using these same structural guidelines?