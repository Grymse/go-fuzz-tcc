# Generated by Grammarinator 23.7

import itertools

from math import inf
from grammarinator.runtime import *

class CLNGenerator(Generator):


    def EOF(self, parent=None):
        pass
    EOF.min_depth = 0

    def program(self, parent=None):
        with RuleContext(self, UnparserRule(name='program', parent=parent)) as current:
            if self._max_depth >= 1:
                for _ in self._model.quantify(current, 0, min=0, max=1):
                    self.global_decl(parent=current)
            UnlexerRule(src='int', parent=current)
            UnlexerRule(src='main', parent=current)
            UnlexerRule(src='(', parent=current)
            UnlexerRule(src=')', parent=current)
            UnlexerRule(src='{', parent=current)
            self.variable_decl_as(parent=current)
            if self._max_depth >= 2:
                for _ in self._model.quantify(current, 1, min=0, max=inf):
                    self.statement(parent=current)
            UnlexerRule(src='return', parent=current)
            UnlexerRule(src='1;', parent=current)
            UnlexerRule(src='}', parent=current)
            return current
    program.min_depth = 3

    def global_decl(self, parent=None):
        with RuleContext(self, UnparserRule(name='global_decl', parent=parent)) as current:
            with AlternationContext(self, [3, 0], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.variable_decl_as(parent=current)
                elif choice0 == 1:
                    if self._max_depth >= 3:
                        for _ in self._model.quantify(current, 0, min=0, max=inf):
                            self.func_decl(parent=current)
            return current
    global_decl.min_depth = 0

    def func(self, parent=None):
        with RuleContext(self, UnparserRule(name='func', parent=parent)) as current:
            self.FUNC(parent=current)
            UnlexerRule(src='(', parent=current)
            if self._max_depth >= 11:
                for _ in self._model.quantify(current, 0, min=0, max=1):
                    self.expr(parent=current)
                    if self._max_depth >= 11:
                        for _ in self._model.quantify(current, 1, min=0, max=inf):
                            UnlexerRule(src=',', parent=current)
                            self.expr(parent=current)
            UnlexerRule(src=')', parent=current)
            return current
    func.min_depth = 2

    def func_decl(self, parent=None):
        with RuleContext(self, UnparserRule(name='func_decl', parent=parent)) as current:
            self.type_specifier(parent=current)
            self.FUNC_DECL(parent=current)
            UnlexerRule(src='(', parent=current)
            if self._max_depth >= 3:
                for _ in self._model.quantify(current, 0, min=0, max=1):
                    self.param_list(parent=current)
            UnlexerRule(src=')', parent=current)
            self.compound_statement(parent=current)
            return current
    func_decl.min_depth = 2

    def param_list(self, parent=None):
        with RuleContext(self, UnparserRule(name='param_list', parent=parent)) as current:
            self.param(parent=current)
            if self._max_depth >= 2:
                for _ in self._model.quantify(current, 0, min=0, max=inf):
                    UnlexerRule(src=',', parent=current)
                    self.param(parent=current)
            return current
    param_list.min_depth = 2

    def param(self, parent=None):
        with RuleContext(self, UnparserRule(name='param', parent=parent)) as current:
            self.type_specifier(parent=current)
            self.ID(parent=current)
            return current
    param.min_depth = 1

    def compound_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='compound_statement', parent=parent)) as current:
            with AlternationContext(self, [0, 0], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='{', parent=current)
                    if self._max_depth >= 5:
                        for _ in self._model.quantify(current, 0, min=0, max=inf):
                            self.var_decl(parent=current)
                    if self._max_depth >= 2:
                        for _ in self._model.quantify(current, 1, min=0, max=inf):
                            self.statement(parent=current)
                    UnlexerRule(src='}', parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='{}', parent=current)
            return current
    compound_statement.min_depth = 0

    def var_decl(self, parent=None):
        with RuleContext(self, UnparserRule(name='var_decl', parent=parent)) as current:
            self.type_specifier(parent=current)
            self.var_decl_list(parent=current)
            UnlexerRule(src=';', parent=current)
            return current
    var_decl.min_depth = 4

    def type_specifier(self, parent=None):
        with RuleContext(self, UnparserRule(name='type_specifier', parent=parent)) as current:
            with AlternationContext(self, [1, 0], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                src = [None, 'char'][choice0]
                rule = [self.number_types, None][choice0]
                if src is not None:
                    UnlexerRule(src=src, parent=current)
                else:
                    rule(parent=current)
            return current
    type_specifier.min_depth = 0

    def number_types(self, parent=None):
        with RuleContext(self, UnparserRule(name='number_types', parent=parent)) as current:
            with AlternationContext(self, [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0], [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                UnlexerRule(src=['int', 'signed int', 'unsigned int', 'short', 'signed short', 'unsigned short', 'long', 'signed long', 'unsigned long', 'float', 'double'][choice0], parent=current)
            return current
    number_types.min_depth = 0

    def arr_decl(self, parent=None):
        with RuleContext(self, UnparserRule(name='arr_decl', parent=parent)) as current:
            with AlternationContext(self, [2, 2], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='const', parent=current)
                    self.number_types(parent=current)
                    self.ID_DECL_ARR_C(parent=current)
                    self.arr_decl_array(parent=current)
                elif choice0 == 1:
                    self.number_types(parent=current)
                    self.ID_DECL_ARR(parent=current)
                    self.arr_decl_array(parent=current)
            return current
    arr_decl.min_depth = 2

    def arr_decl_array(self, parent=None):
        with RuleContext(self, UnparserRule(name='arr_decl_array', parent=parent)) as current:
            with AlternationContext(self, [1, 3, 1, 0, 3, 0], [1, 1, 1, 1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='[', parent=current)
                    self.INT(parent=current)
                    UnlexerRule(src=']', parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='[', parent=current)
                    self.INT(parent=current)
                    UnlexerRule(src=']', parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.comma_separated_value(parent=current)
                elif choice0 == 2:
                    UnlexerRule(src='[', parent=current)
                    self.INT(parent=current)
                    UnlexerRule(src=']', parent=current)
                    UnlexerRule(src='=', parent=current)
                    UnlexerRule(src='{}', parent=current)
                elif choice0 == 3:
                    UnlexerRule(src='[]', parent=current)
                elif choice0 == 4:
                    UnlexerRule(src='[]', parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.comma_separated_value(parent=current)
                elif choice0 == 5:
                    UnlexerRule(src='[]', parent=current)
                    UnlexerRule(src='=', parent=current)
                    UnlexerRule(src='{}', parent=current)
            return current
    arr_decl_array.min_depth = 0

    def string_decl(self, parent=None):
        with RuleContext(self, UnparserRule(name='string_decl', parent=parent)) as current:
            with AlternationContext(self, [2, 2], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='const', parent=current)
                    UnlexerRule(src='char', parent=current)
                    self.ID_DECL_ARR_C(parent=current)
                    UnlexerRule(src='[]', parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.LOREM(parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='char', parent=current)
                    self.ID_DECL_ARR(parent=current)
                    UnlexerRule(src='[]', parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.LOREM(parent=current)
            return current
    string_decl.min_depth = 2

    def comma_separated_value(self, parent=None):
        with RuleContext(self, UnparserRule(name='comma_separated_value', parent=parent)) as current:
            with AlternationContext(self, [3, 2], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.value(parent=current)
                    UnlexerRule(src=',', parent=current)
                    self.comma_separated_value(parent=current)
                elif choice0 == 1:
                    self.value(parent=current)
            return current
    comma_separated_value.min_depth = 2

    def var_decl_list(self, parent=None):
        with RuleContext(self, UnparserRule(name='var_decl_list', parent=parent)) as current:
            with AlternationContext(self, [3, 4], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.variable_id_as(parent=current)
                elif choice0 == 1:
                    self.variable_id_as(parent=current)
                    UnlexerRule(src=',', parent=current)
                    self.var_decl_list(parent=current)
            return current
    var_decl_list.min_depth = 3

    def variable_id_as(self, parent=None):
        with RuleContext(self, UnparserRule(name='variable_id_as', parent=parent)) as current:
            with AlternationContext(self, [2, 11], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.ID_DECL(parent=current)
                elif choice0 == 1:
                    self.ID_DECL(parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.expr(parent=current)
            return current
    variable_id_as.min_depth = 2

    def variable_decl_as(self, parent=None):
        with RuleContext(self, UnparserRule(name='variable_decl_as', parent=parent)) as current:
            with AlternationContext(self, [2, 2, 2, 3, 3], [1, 1, 1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.type_specifier(parent=current)
                    self.ID_DECL(parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.value(parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 1:
                    self.type_specifier(parent=current)
                    self.ID_DECL(parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 2:
                    UnlexerRule(src='const', parent=current)
                    self.type_specifier(parent=current)
                    self.ID_DECL_C(parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.value(parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 3:
                    self.string_decl(parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 4:
                    self.arr_decl(parent=current)
                    UnlexerRule(src=';', parent=current)
            return current
    variable_decl_as.min_depth = 2

    def statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='statement', parent=parent)) as current:
            with AlternationContext(self, [12, 12, 11, 1, 12, 12, 2, 2, 3, 1], [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.switch_statement(parent=current)
                elif choice0 == 1:
                    self.do_while_statement(parent=current)
                elif choice0 == 2:
                    self.for_statement(parent=current)
                elif choice0 == 3:
                    self.compound_statement(parent=current)
                elif choice0 == 4:
                    self.cond_statement(parent=current)
                elif choice0 == 5:
                    self.while_statement(parent=current)
                elif choice0 == 6:
                    self.multiline_comment(parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 7:
                    UnlexerRule(src='//', parent=current)
                    self.LOREM(parent=current)
                    self.statement(parent=current)
                elif choice0 == 8:
                    self.func(parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 9:
                    UnlexerRule(src='return', parent=current)
                    self.ID(parent=current)
                    UnlexerRule(src=';', parent=current)
            return current
    statement.min_depth = 1

    def multiline_comment(self, parent=None):
        with RuleContext(self, UnparserRule(name='multiline_comment', parent=parent)) as current:
            with AlternationContext(self, [1, 1], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='/*', parent=current)
                    self.LOREM(parent=current)
                    UnlexerRule(src='*/', parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='/*', parent=current)
                    UnlexerRule(src='\n', parent=current)
                    UnlexerRule(src='*', parent=current)
                    self.LOREM(parent=current)
                    UnlexerRule(src='\n', parent=current)
                    UnlexerRule(src='*/', parent=current)
            return current
    multiline_comment.min_depth = 1

    def switch_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='switch_statement', parent=parent)) as current:
            with AlternationContext(self, [11, 11], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='switch', parent=current)
                    UnlexerRule(src='(', parent=current)
                    self.expr(parent=current)
                    UnlexerRule(src=')', parent=current)
                    UnlexerRule(src='{', parent=current)
                    if self._max_depth >= 3:
                        for _ in self._model.quantify(current, 0, min=0, max=inf):
                            self.case_statement(parent=current)
                    UnlexerRule(src='}', parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='switch', parent=current)
                    UnlexerRule(src='(', parent=current)
                    self.expr(parent=current)
                    UnlexerRule(src=')', parent=current)
                    UnlexerRule(src='{', parent=current)
                    if self._max_depth >= 3:
                        for _ in self._model.quantify(current, 1, min=0, max=inf):
                            self.case_statement(parent=current)
                    UnlexerRule(src='default', parent=current)
                    UnlexerRule(src=':', parent=current)
                    self.statement(parent=current)
                    UnlexerRule(src='}', parent=current)
            return current
    switch_statement.min_depth = 11

    def case_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='case_statement', parent=parent)) as current:
            with AlternationContext(self, [2, 2, 2, 2], [1, 1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='case', parent=current)
                    self.INT(parent=current)
                    UnlexerRule(src=':', parent=current)
                    self.statement(parent=current)
                    UnlexerRule(src='break', parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='case', parent=current)
                    self.CHAR(parent=current)
                    UnlexerRule(src=':', parent=current)
                    self.statement(parent=current)
                    UnlexerRule(src='break', parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 2:
                    UnlexerRule(src='case', parent=current)
                    self.INT(parent=current)
                    UnlexerRule(src=':', parent=current)
                    self.statement(parent=current)
                elif choice0 == 3:
                    UnlexerRule(src='case', parent=current)
                    self.CHAR(parent=current)
                    UnlexerRule(src=':', parent=current)
                    self.statement(parent=current)
            return current
    case_statement.min_depth = 2

    def do_while_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='do_while_statement', parent=parent)) as current:
            UnlexerRule(src='do', parent=current)
            self.loop_statement(parent=current)
            UnlexerRule(src='while', parent=current)
            UnlexerRule(src='(', parent=current)
            self.expr(parent=current)
            UnlexerRule(src=')', parent=current)
            UnlexerRule(src=';', parent=current)
            return current
    do_while_statement.min_depth = 11

    def while_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='while_statement', parent=parent)) as current:
            UnlexerRule(src='while', parent=current)
            UnlexerRule(src='(', parent=current)
            self.expr(parent=current)
            UnlexerRule(src=')', parent=current)
            self.loop_statement(parent=current)
            return current
    while_statement.min_depth = 11

    def for_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='for_statement', parent=parent)) as current:
            with AlternationContext(self, [11, 10], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='for', parent=current)
                    UnlexerRule(src='(', parent=current)
                    UnlexerRule(src='int', parent=current)
                    self.ID_DECL(parent=current)
                    UnlexerRule(src='=', parent=current)
                    UnlexerRule(src='0', parent=current)
                    UnlexerRule(src=';', parent=current)
                    self.condition(parent=current)
                    UnlexerRule(src=';', parent=current)
                    self.ID_AS(parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.expr(parent=current)
                    UnlexerRule(src=')', parent=current)
                    self.loop_statement(parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='for', parent=current)
                    UnlexerRule(src='(', parent=current)
                    UnlexerRule(src=';', parent=current)
                    self.condition(parent=current)
                    UnlexerRule(src=';', parent=current)
                    UnlexerRule(src=')', parent=current)
                    self.loop_statement(parent=current)
            return current
    for_statement.min_depth = 10

    def loop_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='loop_statement', parent=parent)) as current:
            with AlternationContext(self, [0, 0, 1], [1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='{', parent=current)
                    if self._max_depth >= 1:
                        for _ in self._model.quantify(current, 0, min=0, max=inf):
                            self.loop_inner_statement(parent=current)
                    UnlexerRule(src='}', parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='{}', parent=current)
                elif choice0 == 2:
                    self.loop_inner_statement(parent=current)
            return current
    loop_statement.min_depth = 0

    def loop_inner_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='loop_inner_statement', parent=parent)) as current:
            with AlternationContext(self, [2, 0, 0], [1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.statement(parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='break', parent=current)
                    UnlexerRule(src=';', parent=current)
                elif choice0 == 2:
                    UnlexerRule(src='continue', parent=current)
                    UnlexerRule(src=';', parent=current)
            return current
    loop_inner_statement.min_depth = 0

    def cond_statement(self, parent=None):
        with RuleContext(self, UnparserRule(name='cond_statement', parent=parent)) as current:
            with AlternationContext(self, [11, 11], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='if', parent=current)
                    UnlexerRule(src='(', parent=current)
                    self.expr(parent=current)
                    UnlexerRule(src=')', parent=current)
                    self.statement(parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='if', parent=current)
                    UnlexerRule(src='(', parent=current)
                    self.expr(parent=current)
                    UnlexerRule(src=')', parent=current)
                    self.statement(parent=current)
                    UnlexerRule(src='else', parent=current)
                    self.statement(parent=current)
            return current
    cond_statement.min_depth = 11

    def expr(self, parent=None):
        with RuleContext(self, UnparserRule(name='expr', parent=parent)) as current:
            with AlternationContext(self, [11, 10], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.ID_AS(parent=current)
                    UnlexerRule(src='=', parent=current)
                    self.expr(parent=current)
                elif choice0 == 1:
                    self.condition(parent=current)
            return current
    expr.min_depth = 10

    def condition(self, parent=None):
        with RuleContext(self, UnparserRule(name='condition', parent=parent)) as current:
            with AlternationContext(self, [9, 11], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.disjunction(parent=current)
                elif choice0 == 1:
                    self.disjunction(parent=current)
                    UnlexerRule(src='?', parent=current)
                    self.expr(parent=current)
                    UnlexerRule(src=':', parent=current)
                    self.condition(parent=current)
            return current
    condition.min_depth = 9

    def disjunction(self, parent=None):
        with RuleContext(self, UnparserRule(name='disjunction', parent=parent)) as current:
            with AlternationContext(self, [8, 9], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.conjunction(parent=current)
                elif choice0 == 1:
                    self.disjunction(parent=current)
                    UnlexerRule(src='||', parent=current)
                    self.conjunction(parent=current)
            return current
    disjunction.min_depth = 8

    def conjunction(self, parent=None):
        with RuleContext(self, UnparserRule(name='conjunction', parent=parent)) as current:
            with AlternationContext(self, [7, 8], [1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.comparison(parent=current)
                elif choice0 == 1:
                    self.conjunction(parent=current)
                    UnlexerRule(src='&&', parent=current)
                    self.comparison(parent=current)
            return current
    conjunction.min_depth = 7

    def comparison(self, parent=None):
        with RuleContext(self, UnparserRule(name='comparison', parent=parent)) as current:
            with AlternationContext(self, [6, 6, 6], [1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.relation(parent=current)
                elif choice0 == 1:
                    self.relation(parent=current)
                    UnlexerRule(src='==', parent=current)
                    self.relation(parent=current)
                elif choice0 == 2:
                    self.relation(parent=current)
                    UnlexerRule(src='!=', parent=current)
                    self.relation(parent=current)
            return current
    comparison.min_depth = 6

    def relation(self, parent=None):
        with RuleContext(self, UnparserRule(name='relation', parent=parent)) as current:
            with AlternationContext(self, [5, 5, 5, 5, 5], [1, 1, 1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.sum_(parent=current)
                elif choice0 == 1:
                    self.sum_(parent=current)
                    UnlexerRule(src='<=', parent=current)
                    self.sum_(parent=current)
                elif choice0 == 2:
                    self.sum_(parent=current)
                    UnlexerRule(src='<', parent=current)
                    self.sum_(parent=current)
                elif choice0 == 3:
                    self.sum_(parent=current)
                    UnlexerRule(src='>=', parent=current)
                    self.sum_(parent=current)
                elif choice0 == 4:
                    self.sum_(parent=current)
                    UnlexerRule(src='>', parent=current)
                    self.sum_(parent=current)
            return current
    relation.min_depth = 5

    def sum_(self, parent=None):
        with RuleContext(self, UnparserRule(name='sum_', parent=parent)) as current:
            with AlternationContext(self, [4, 5, 5], [1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.term(parent=current)
                elif choice0 == 1:
                    self.sum_(parent=current)
                    UnlexerRule(src='+', parent=current)
                    self.term(parent=current)
                elif choice0 == 2:
                    self.sum_(parent=current)
                    UnlexerRule(src='-', parent=current)
                    self.term(parent=current)
            return current
    sum_.min_depth = 4

    def term(self, parent=None):
        with RuleContext(self, UnparserRule(name='term', parent=parent)) as current:
            with AlternationContext(self, [3, 4, 4], [1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.factor(parent=current)
                elif choice0 == 1:
                    self.term(parent=current)
                    UnlexerRule(src='*', parent=current)
                    self.factor(parent=current)
                elif choice0 == 2:
                    self.term(parent=current)
                    UnlexerRule(src='/', parent=current)
                    self.factor(parent=current)
            return current
    term.min_depth = 3

    def factor(self, parent=None):
        with RuleContext(self, UnparserRule(name='factor', parent=parent)) as current:
            with AlternationContext(self, [3, 3, 2], [1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    UnlexerRule(src='!', parent=current)
                    self.factor(parent=current)
                elif choice0 == 1:
                    UnlexerRule(src='-', parent=current)
                    self.factor(parent=current)
                elif choice0 == 2:
                    self.primary(parent=current)
            return current
    factor.min_depth = 2

    def primary(self, parent=None):
        with RuleContext(self, UnparserRule(name='primary', parent=parent)) as current:
            with AlternationContext(self, [2, 2, 1, 2, 3, 3, 11], [1, 1, 1, 1, 1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.value(parent=current)
                elif choice0 == 1:
                    self.value(parent=current)
                    self.multiline_comment(parent=current)
                elif choice0 == 2:
                    self.ID(parent=current)
                elif choice0 == 3:
                    self.ID(parent=current)
                    self.multiline_comment(parent=current)
                elif choice0 == 4:
                    self.func(parent=current)
                    self.multiline_comment(parent=current)
                elif choice0 == 5:
                    self.func(parent=current)
                elif choice0 == 6:
                    UnlexerRule(src='(', parent=current)
                    self.expr(parent=current)
                    UnlexerRule(src=')', parent=current)
            return current
    primary.min_depth = 1

    def value(self, parent=None):
        with RuleContext(self, UnparserRule(name='value', parent=parent)) as current:
            with AlternationContext(self, [1, 1, 1], [1, 1, 1]) as weights0:
                choice0 = self._model.choice(current, 0, weights0)
                if choice0 == 0:
                    self.INT(parent=current)
                elif choice0 == 1:
                    self.INT(parent=current)
                    UnlexerRule(src='.', parent=current)
                    self.INT(parent=current)
                    UnlexerRule(src='f', parent=current)
                elif choice0 == 2:
                    self.CHAR(parent=current)
            return current
    value.min_depth = 1

    def INT(self, parent=None):
        with RuleContext(self, UnlexerRule(name='INT', parent=parent)) as current:
            if self._max_depth >= 0:
                for _ in self._model.quantify(current, 0, min=1, max=inf):
                    UnlexerRule(src=self._model.charset(current, 0, self._charsets[1]), parent=current)
            return current
    INT.min_depth = 0

    def CHAR(self, parent=None):
        with RuleContext(self, UnlexerRule(name='CHAR', parent=parent)) as current:
            UnlexerRule(src='\'', parent=current)
            UnlexerRule(src=self._model.charset(current, 0, self._charsets[2]), parent=current)
            UnlexerRule(src='\'', parent=current)
            return current
    CHAR.min_depth = 0

    def ID(self, parent=None):
        with RuleContext(self, UnlexerRule(name='ID', parent=parent)) as current:
            UnlexerRule(src=self._model.charset(current, 0, self._charsets[3]), parent=current)
            if self._max_depth >= 0:
                for _ in self._model.quantify(current, 0, min=0, max=inf):
                    UnlexerRule(src=self._model.charset(current, 1, self._charsets[4]), parent=current)
            return current
    ID.min_depth = 0

    def ID_DECL(self, parent=None):
        with RuleContext(self, UnlexerRule(name='ID_DECL', parent=parent)) as current:
            self.ID(parent=current)
            return current
    ID_DECL.min_depth = 1

    def ID_AS(self, parent=None):
        with RuleContext(self, UnlexerRule(name='ID_AS', parent=parent)) as current:
            self.ID(parent=current)
            return current
    ID_AS.min_depth = 1

    def ID_DECL_ARR(self, parent=None):
        with RuleContext(self, UnlexerRule(name='ID_DECL_ARR', parent=parent)) as current:
            self.ID(parent=current)
            return current
    ID_DECL_ARR.min_depth = 1

    def ID_DECL_ARR_C(self, parent=None):
        with RuleContext(self, UnlexerRule(name='ID_DECL_ARR_C', parent=parent)) as current:
            self.ID(parent=current)
            return current
    ID_DECL_ARR_C.min_depth = 1

    def ID_DECL_C(self, parent=None):
        with RuleContext(self, UnlexerRule(name='ID_DECL_C', parent=parent)) as current:
            self.ID(parent=current)
            return current
    ID_DECL_C.min_depth = 1

    def FUNC_DECL(self, parent=None):
        with RuleContext(self, UnlexerRule(name='FUNC_DECL', parent=parent)) as current:
            self.ID(parent=current)
            return current
    FUNC_DECL.min_depth = 1

    def FUNC(self, parent=None):
        with RuleContext(self, UnlexerRule(name='FUNC', parent=parent)) as current:
            self.ID(parent=current)
            return current
    FUNC.min_depth = 1

    def LOREM(self, parent=None):
        with RuleContext(self, UnlexerRule(name='LOREM', parent=parent)) as current:
            UnlexerRule(src='"', parent=current)
            if self._max_depth >= 0:
                for _ in self._model.quantify(current, 0, min=0, max=inf):
                    UnlexerRule(src=self._model.charset(current, 0, self._charsets[0]), parent=current)
            UnlexerRule(src='"', parent=current)
            return current
    LOREM.min_depth = 0

    def WS(self, parent=None):
        with RuleContext(self, UnlexerRule(name='WS', parent=parent)) as current:
            if self._max_depth >= 0:
                for _ in self._model.quantify(current, 0, min=1, max=inf):
                    UnlexerRule(src=self._model.charset(current, 0, self._charsets[5]), parent=current)
            return current
    WS.min_depth = 0

    _default_rule = program

    _charsets = {
        0: list(itertools.chain.from_iterable([range(32, 127)])),
        1: list(itertools.chain.from_iterable([range(48, 58)])),
        2: list(itertools.chain.from_iterable([range(48, 58), range(65, 91), range(97, 123)])),
        3: list(itertools.chain.from_iterable([range(65, 91), range(95, 96), range(97, 123)])),
        4: list(itertools.chain.from_iterable([range(48, 58), range(65, 91), range(95, 96), range(97, 123)])),
        5: list(itertools.chain.from_iterable([range(9, 10), range(10, 11), range(13, 14), range(32, 33)])),
    }
